name = "CodeIntel"
owner = "ActiveState"
namespace = "github.com/ActiveState/CodeIntel"
version = "master"
environments = "dev,qa,prod"

platform "Linux64Label" {
  os = "linux"
  architecture = "amd64"
  libc = "glibc-2.25"
  compiler = "gcc-7"
}

platform "Windows10Label" {
  os = "windows"
  version = "10"
}

language "Go" {
  version = "1.10"
  constraint {
    platform = "Windows10Label,Linux64Label"
    environment = "dev,qa,prod"
  }
  package "golang.org/x/crypto" {
    version = "*"
    build {
      debug = "$variable.DEBUG"
    }
  }
  package "gopkg.in/yaml.v2" {
    version = "2.*"
    build {
      override = "--foo --bar --debug $variable.DEBUG --libDir $variable.PYTHONPATH --libc $platform.libc"
    }
  }
}

language "Python" {
  version = "2.7.12"
  constraint {
    platform = "Windows10Label,Linux64Label"
  }
  package "apsw" {
    version = "3.8.11.1"
    build {
      debug = "$variable.DEBUG"
    }
  }
  package "peewee" {
    version = "2.9.1"
    build {
      override = "--foo --bar --debug $variable.DEBUG"
    }
  }
}

variable "DEBUG" {
  value = "true"
}

variable "PYTHONPATH" {
  value = "%projectDir%/src:%projectDir%/tests"
  constraint {
    environment = "dev,qa"
  }
}

variable "PYTHONPATH" {
  value = "%projectDir%/src:%projectDir%/tests"
}

hook {
  name = "FIRST_INSTALL"
  value = "%pythonExe% %projectDir%/setup.py prepare"
}

hook {
  name = "AFTER_UPDATE"
  value = "%pythonExe% %projectDir%/setup.py prepare"
}

command "tests" {
  value = "pytest %projectDir%/tests"
}

command "debug" {
  value = "debug foo"
}
