package main

//--------------------------------------------------------------------------------------------------
//
// This program will
//
//     - Compile the Python code using qmake.
//     - Compile the web assets using yarn.
//     - Create a tarball and put in ~/.toolbox-tarballs
//
// The toolbox build process will copy the tarball from ~/.toolbox-tarballs into the release
// tarball.
//
// This program assumes that 'qmake' is installed and in the search path.  On Darwin, we install
// QT using 'brew install qt'.
//
// This program also assumes that it is running in a Python >= 3.4 virtual environment.
//
//--------------------------------------------------------------------------------------------------

import (
  "fmt"
  "os"
  "os/exec"
  "path"
  "path/filepath"
)

type Builder struct {
}

func (this Builder) build() {
  repoDir, _ :=  filepath.Abs(filepath.Dir(os.Args[0]))
  runtimeDir := path.Join(repoDir, "runtime")

  // Find the python executable
  pythonFilepath, err := exec.LookPath("python")
  if err != nil {
    fmt.Printf("Could not find python executable. %s\n", err.Error())
    os.Exit(1);
  }
  fmt.Printf("using python: '%s'\n", pythonFilepath);

  // Find the qmake executable
  qmakeFilepath, err := exec.LookPath("qmake")
  if err != nil {
    fmt.Printf("Could not find qmake executable. %s\n", err.Error())
    os.Exit(1);
  }
  fmt.Printf("using qmake: '%s'\n", qmakeFilepath);

  // Find the yarn executable
  yarnFilepath, err := exec.LookPath("yarn")
  if err != nil {
    fmt.Printf("Could not find yarn executable. %s\n", err.Error())
    os.Exit(1);
  }
  fmt.Printf("using yarn: '%s'\n", yarnFilepath);

  // The PgAdmin process needs to know where to find Python
  pythonDir := filepath.Dir(filepath.Dir(pythonFilepath))
  fmt.Printf("pythonDir: %s\n", pythonDir)
  os.Setenv("PGADMIN_PYTHON_DIR", pythonDir)

  fmt.Printf("env: %s\n", os.Getenv("PGADMIN_PYTHON_DIR"))

  // Exec 'qmake'
  command := "qmake"
  os.Chdir(runtimeDir)
  output, err := exec.Command(command).Output()
  os.Chdir(repoDir)
  if err != nil {
    fmt.Printf("Failed to execute '%s'. %s %s\n", command, output, err.Error())
    os.Exit(1)
  }
  
  // // Exec 'qmake'
  // command := fmt.Sprintf("%s", qmakeFilepath)
  // fmt.Printf("--- qmake\n");
  // os.Chdir(runtimeDir)
  // output, err := exec.Command("bash", "-c", command).Output()
  // os.Chdir(repoDir)
  // if err != nil {
  //   fmt.Printf("Failed to execute command '%s'. %s %s\n", command, output, err.Error())
  //   os.Exit(1)
  // }

  // // Exec 'qmake CONFIG+=release'
  // command = fmt.Sprintf("%s CONFIG+=release", qmakeFilepath)
  // fmt.Printf("--- qmake CONFIG+=release\n");
  // os.Chdir(runtimeDir)
  // output, err = exec.Command("bash", "-c", command).Output()
  // os.Chdir(repoDir)
  // if err != nil {
  //   fmt.Printf("Failed to execute command '%s'. %s %s\n", command, output, err.Error())
  //   os.Exit(1)
  // }

  // // Build the web assets
  // webDir := path.Join(repoDir, "pgadmin", "web")

  // // Exec 'yarn install'
  // command = fmt.Sprintf("yarn install")
  // fmt.Printf("--- yarn install\n");
  // os.Chdir(webDir)
  // output, err = exec.Command("bash", "-c", command).Output()
  // os.Chdir(repoDir)
  // if err != nil {
  //   fmt.Printf("Failed to execute command '%s'. %s %s\n", command, output, err.Error())
  //   os.Exit(1)
  // }

  // // Exec 'yarn run build'
  // command = fmt.Sprintf("yarn run bundle")
  // fmt.Printf("--- yarn run bundle\n");
  // os.Chdir(webDir)
  // output, err = exec.Command("bash", "-c", command).Output()
  // os.Chdir(repoDir)
  // if err != nil {
  //   fmt.Printf("Failed to execute command '%s'. %s\n", command, output, err.Error())
  //   os.Exit(1)
  // }

  // // Tar the pgadmin directory
  // tarCommand := fmt.Sprintf("tar -C  %s -czf %s pgadmin", repoDir, this.TarballFilepath)
  // context.Logf("--- %s", tarCommand);
  // _, err = exec.Command("bash", "-c", tarCommand).Output()
  // if err != nil {
  //   return fmt.Errorf("error tarring pgadmin. %s\n", err.Error())
  // }

  
}

func main() {
  var builder Builder
  builder.build()
}
