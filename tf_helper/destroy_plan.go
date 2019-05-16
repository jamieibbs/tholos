package tf_helper

import (
    "fmt"
    "log"

    "github.com/mhlias/tholos/tholos"
)

func (c *Config) DestroyPlan(tholos_conf *tholos.Tholos_config, parallelism int16) {

    cmd_name := "rm"

    exec_args := []string{"-rf", ".terraform"}

    log.Println("[INFO] Deleting Terraform cache directory.")

    if !ExecCmd(cmd_name, exec_args) {
        log.Fatal("[ERROR] Failed to delete Terraform cache directory. Aborting.")
    }

    cmd_name = "rm"

    exec_args = []string{"-f", "plans/destroyplan.tfplan"}

    log.Println("[INFO] Deleting Terraform old destroy plan.")

    if !ExecCmd(cmd_name, exec_args) {
        log.Fatal("[ERROR] Failed to delete Terraform old destroy plan. Aborting.")
    }

    cmd_name = "terraform"

    exec_args = []string{"get", "-update=true"}

    log.Println("[INFO] Fetching Terraform modules and updating existing ones.")

    if !ExecCmd(cmd_name, exec_args) {
        log.Fatal("[ERROR] Failed to fetch Terraform modules. Aborting.")
    }

    c.Setup_remote_state()

    exec_args = []string{"plan -destroy", fmt.Sprintf("-parallelism=%d", parallelism), "-module-depth=3", "-refresh=true", "-out=plans/destroyplan.tfplan", "-var-file=params/env.tfvars"}

    if len(c.TargetsTF) > 0 {
        for _, t := range c.TargetsTF {
            exec_args = append(exec_args, fmt.Sprintf("-target=%s", t))
        }
    }

    log.Println("[INFO] Running Terraform destroy plan.")

    if !ExecCmd(cmd_name, exec_args) {
        log.Fatal("[ERROR] Failed to run Terraform destroy plan. Aborting.")
    }

}
