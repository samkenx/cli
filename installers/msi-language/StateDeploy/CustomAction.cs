using Microsoft.Deployment.WindowsInstaller;
using System;
using System.Text;
using System.IO;
using System.Net;
using System.Collections.ObjectModel;
using System.Windows.Forms;
using System.Linq;
using System.Web.Script.Serialization;
using System.Collections.Generic;
using Newtonsoft.Json;
using System.Security.Cryptography;
using System.IO.Compression;

namespace StateDeploy
{

    enum Shell
    {
        Cmd,
        Powershell,
    }

    public class CustomActions
    {

        private class VersionInfo
        {
            public string version;
            public string sha256v2;
        }

        public static ActionResult InstallStateTool(Session session, out string stateToolPath)
        {
            ActiveState.RollbarHelper.ConfigureRollbarSingleton();
            stateToolPath = "";

            string stateURL = "https://s3.ca-central-1.amazonaws.com/cli-update/update/state/unstable/";
            string windowsJSON = "windows-amd64.json";
            string jsonURL = stateURL + windowsJSON;
            string tempDir = Path.Combine(Path.GetTempPath(), "ActiveState");
            ServicePointManager.SecurityProtocol |= SecurityProtocolType.Tls11 | SecurityProtocolType.Tls12;

            string versionInfoString;
            session.Log(string.Format("Downloading JSON with URL: {0}", jsonURL));
            try
            {
                WebClient client = new WebClient();
                versionInfoString = client.DownloadString(jsonURL);
            }
            catch (WebException e)
            {
                string msg = string.Format("Encoutered exception downloading state tool json info file: {0}", e.ToString());
                session.Log(msg);
                ActiveState.RollbarHelper.Report(msg);
                return ActionResult.Failure;
            }

            VersionInfo info = JsonConvert.DeserializeObject<VersionInfo>(versionInfoString);

            string windowsZip = "windows-amd64.zip";
            string zipURL = stateURL + info.version + "/" + windowsZip;
            string zipPath = Path.Combine(tempDir, windowsZip);
            session.Log(string.Format("Downloading zip file with URL: {0}", zipURL));
            try
            {
                WebClient client = new WebClient();
                client.DownloadFile(zipURL, zipPath);
            }
            catch (WebException e)
            {
                string msg = string.Format("Encoutered exception downloading state tool zip file: {0}", e.ToString());
                session.Log(msg);
                ActiveState.RollbarHelper.Report(msg);
                return ActionResult.Failure;
            }

            SHA256 sha = SHA256.Create();
            FileStream fInfo = File.OpenRead(zipPath);
            string zipHash = BitConverter.ToString(sha.ComputeHash(fInfo)).Replace("-", string.Empty).ToLower();
            if (zipHash != info.sha256v2)
            {
                string msg = string.Format("SHA256 checksum did not match, expected: {0} actual: {1}", info.sha256v2, zipHash.ToString());
                session.Log(msg);
                ActiveState.RollbarHelper.Report(msg);
                return ActionResult.Failure;
            }

            try
            {
                ZipFile.ExtractToDirectory(zipPath, tempDir);
            }
            catch (Exception e)
            {
                string msg = string.Format("Could not extract State Tool, encountered exception: {0)", e);
                session.Log(msg);
                ActiveState.RollbarHelper.Report(msg);
                return ActionResult.Failure;
            }

            stateToolPath = Path.Combine(Environment.GetFolderPath(Environment.SpecialFolder.ApplicationData), "ActiveState", "bin", "state.exe");
            try
            {
                File.Move(Path.Combine(tempDir, "windows-amd64.exe"), stateToolPath);
            }
            catch (Exception e)
            {
                string msg = string.Format("Could not move State Tool executable, encountered exception: {0}", e);
                session.Log(msg);
                ActiveState.RollbarHelper.Report(msg);
                return ActionResult.Failure;
            }

            return ActionResult.Success;
        }

        private static ActionResult Login(Session session, string stateToolPath)
        {
            string username = session.CustomActionData["AS_USERNAME"];
            string password = session.CustomActionData["AS_PASSWORD"];
            string totp = session.CustomActionData["AS_TOTP"];

            if (username == "" && password == "" && totp == "")
            {
                session.Log("No login information provided, not executing login");
                return ActionResult.Success;
            }

            string authCmd;
            if (totp != "")
            {
                session.Log("Attempting to log in with TOTP token");
                authCmd = stateToolPath + " auth" + " --totp " + totp;
            } else
            {
                session.Log(string.Format("Attempting to login as user: {0}", username));
                authCmd = stateToolPath + " auth" + " --username " + username + " --password " + password;
            }

            string output;
            Status.ProgressBar.StatusMessage(session, "Authenticating...");
            ActionResult runResult = ActiveState.Command.Run(session, authCmd, ActiveState.Shell.Cmd, out output);
            if (runResult.Equals(ActionResult.UserExit))
            {
                // Catch cancel and return
                return runResult;
            }
            else if (runResult == ActionResult.Failure)
            {
                Record record = new Record();
                session.Log(string.Format("Output: {0}", output));
                var errorOutput = FormatErrorOutput(output);
                record.FormatString = string.Format("Platform login failed with error:\n{0}", errorOutput);

                session.Message(InstallMessage.Error | (InstallMessage)MessageBoxButtons.OK, record);
                return runResult;
            }
            // The auth command did not fail but the username we expected is not present in the output meaning
            // another user is logged into the State Tool 
            else if (!output.Contains(username))
            {
                Record record = new Record();
                var errorOutput = string.Format("Could not log in as {0}, currently logged in as another user. To correct this please start a command prompt and execute `{1} auth logout` and try again", username, stateToolPath);
                record.FormatString = string.Format("Failed with error:\n{0}", errorOutput);

                session.Message(InstallMessage.Error | (InstallMessage)MessageBoxButtons.OK, record);
                return ActionResult.Failure;
            }
            return ActionResult.Success;
        }

        public struct InstallSequenceElement
        {
            public readonly string SubCommand;
            public readonly string Description;

            public InstallSequenceElement(string subCommand, string description)
            {
                this.SubCommand = subCommand;
                this.Description = description;
            }
        };


        [CustomAction]
        public static ActionResult StateDeploy(Session session)
        {
            ActiveState.RollbarHelper.ConfigureRollbarSingleton();

            string stateToolPath;
            ActionResult res = InstallStateTool(session, out stateToolPath);
            if (res != ActionResult.Success) {
                return res;
            }
            session.Log("Starting state deploy with state tool at " + stateToolPath);

            res = Login(session, stateToolPath);
            if (res.Equals(ActionResult.Failure))
            {
                return res;
            }

            Status.ProgressBar.StatusMessage(session, string.Format("Deploying project {0}...", session.CustomActionData["PROJECT_OWNER_AND_NAME"]));
            Status.ProgressBar.StatusMessage(session, string.Format("Preparing deployment of {0}...", session.CustomActionData["PROJECT_OWNER_AND_NAME"]));

            var sequence = new ReadOnlyCollection<InstallSequenceElement>(
                new[]
                {
                    new InstallSequenceElement("install", string.Format("Installing {0}", session.CustomActionData["PROJECT_OWNER_AND_NAME"])),
                    new InstallSequenceElement("configure", "Updating system environment"),
                    new InstallSequenceElement("symlink", "Creating shortcut directory"),
                });

            try
            {
                foreach (var seq in sequence)
                {
                    string deployCmd = BuildDeployCmd(session, seq.SubCommand, stateToolPath);
                    session.Log(string.Format("Executing deploy command: {0}", deployCmd));

                    Status.ProgressBar.Increment(session, 1);
                    Status.ProgressBar.StatusMessage(session, seq.Description);

                    string output;
                    var runResult = ActiveState.Command.Run(session, deployCmd, ActiveState.Shell.Cmd, out output);
                    if (runResult.Equals(ActionResult.UserExit))
                    {
                        // Catch cancel and return
                        return runResult;
                    }
                    else if (runResult == ActionResult.Failure)
                    {
                        Record record = new Record();
                        var errorOutput = FormatErrorOutput(output);
                        record.FormatString = String.Format("{0} failed with error:\n{1}", seq.Description, errorOutput);

                        MessageResult msgRes = session.Message(InstallMessage.Error | (InstallMessage)MessageBoxButtons.OK, record);
                        return runResult;
                    }
                }
            }
            catch (Exception objException)
            {
                session.Log(string.Format("Caught exception: {0}", objException));
                ActiveState.RollbarHelper.Report(string.Format("Caught exception: {0}", objException));
                return ActionResult.Failure;
            }

            Status.ProgressBar.Increment(session, 1);
            return ActionResult.Success;
        }

        /// <summary>
        /// FormatErrorOutput formats the output of a state tool command optimized for display in an error dialog
        /// </summary>
        /// <param name="cmdOutput">
        /// the output from a state tool command run with `--output=json`
        /// </param>
        private static string FormatErrorOutput(string cmdOutput)
        {
            return string.Join("\n", cmdOutput.Split('\x00').Select(blob =>
            {
                try
                {
                    var json = new JavaScriptSerializer();
                    var data = json.Deserialize<Dictionary<string, string>>(blob);
                    var error = data["Error"];
                    return error;
                }
                catch (Exception)
                {
                    return blob;
                }
            }).ToList());

        }

        private static string BuildDeployCmd(Session session, string subCommand, string stateToolPath)
        {
            string installDir = session.CustomActionData["INSTALLDIR"];
            string projectName = session.CustomActionData["PROJECT_OWNER_AND_NAME"];
            string isModify = session.CustomActionData["IS_MODIFY"];

            StringBuilder deployCMDBuilder = new StringBuilder(stateToolPath + " deploy " + subCommand);
            if (isModify == "true")
            {
                deployCMDBuilder.Append(" --force");
            }

            deployCMDBuilder.Append(" --output json");

            // We quote the string here as Windows paths that contain spaces must be quoted.
            // We also account for a path ending with a slash and ensure that the quote character
            // isn't preserved.
            deployCMDBuilder.AppendFormat(" {0} --path=\"{1}\\\"", projectName, @installDir);

            return deployCMDBuilder.ToString();
        }
    }
}
