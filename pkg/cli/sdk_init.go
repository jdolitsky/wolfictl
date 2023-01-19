package cli

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"

	"github.com/openvex/go-vex/pkg/vex"
	"github.com/spf13/cobra"
	"github.com/wolfi-dev/wolfictl/pkg/sdk"
)

func SDKInit() *cobra.Command {
	p := &sdkListParams{}
	cmd := &cobra.Command{
		Use:           "init <path>",
		Short:         "provide a directory to initialize an SDK environment",
		Args:          cobra.ExactArgs(1),
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			targetDir := args[0]
			for _, d := range []string{
				sdk.StaticRootDirectory,
				sdk.StaticConfigsDirectory,
				sdk.StaticMountDirectory,
			} {
				entries, err := sdk.Static.ReadDir(d)
				if err != nil {
					return err
				}
				for _, entry := range entries {
					srcPath := path.Join(d, entry.Name())
					dstPath := path.Join(targetDir, d, entry.Name())
					if entry.IsDir() {
						// TODO: recursive copy of embedded
						fmt.Printf("creating dir %s\n", dstPath)
						if err := os.MkdirAll(dstPath, 0755); err != nil {
							return err
						}
						continue
					}
					fmt.Printf("creating file %s\n", dstPath)
					b, err := sdk.Static.ReadFile(srcPath)
					if err != nil {
						return nil
					}
					mode := fs.FileMode(0644)
					if strings.HasSuffix(dstPath, ".sh") {
						mode = 0777
					}
					if err := os.WriteFile(dstPath, b, mode); err != nil {
						return nil
					}
				}
			}
			return nil
		},
	}
	p.addFlagsTo(cmd)
	return cmd
}

type sdkListParams struct {
	vuln       string
	history    bool
	unresolved bool
}

func (p *sdkListParams) addFlagsTo(cmd *cobra.Command) {
	cmd.Flags().StringVar(&p.vuln, "vuln", "", "vulnerability ID for advisory")
	cmd.Flags().BoolVar(&p.history, "history", false, "show full history for advisories")
	cmd.Flags().BoolVar(&p.unresolved, "unresolved", false, fmt.Sprintf("only show advisories whose latest status is %s or %s", vex.StatusAffected, vex.StatusUnderInvestigation))
}
