using System;
using Microsoft.Deployment.WindowsInstaller;
using ActiveState;
using System.Windows.Forms;

namespace UICheck
{
    public class Check
    {
        [CustomAction]
        public static ActionResult Level(Session session)
        {
            session.Log("Begin UICheck");
            RollbarHelper.ConfigureRollbarSingleton(session["COMMIT_ID"]);

            int level = 0;
            try
            {
                level = int.Parse(session["UILevel"]);
            } catch (Exception e)
            {
                string msg = string.Format("Could not parse UI Level property. Exception: {0}", e.ToString());
                session.Log(msg);
                RollbarReport.Error(msg, session);
                return ActionResult.Failure;
            }

            session.Log("UI level is: {0}", level);
            if (level < 5)
            {
                // Present message to user
                Record record = new Record();
                record.FormatString = string.Format("Installation not suppored for reduced UI Modes. Please run installer again with full UI");

                session.Message(InstallMessage.Error | (InstallMessage)MessageBoxButtons.OK, record);
                return ActionResult.Failure;
            }

            return ActionResult.Success;
        }
    }
}
