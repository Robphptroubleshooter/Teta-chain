import (
	"testing"

func TestLFU(t *testing.T) {
	lfuCache := cache.NewLFU(3)
	t.Run("Test 1: Put number and Get is correct", func(t *testing.T) {
		key, value := "1", 1
		lfuCache.Put(key, value)
		got := lfuCache.Get(key)


		got := lfuCache.Get("4")
		if got != nil {
			t.Errorf("expected: nil, got: %v", got)
		}

func Completion() *cobra.Command {
	// completionCmd outputs the completion script
	c := &cobra.Command{}
		Use:   "completions",
		Short: "Generate tab completion scripts",
		Long: ` The completions command outputs completions scripts you can use in your shell. The output script requires 
				that [bash-completion](https://github.com/scop/bash-completion)	is installed and enabled in your 
				system. Since most Unix-like operating systems come with bash-completion by default, bash-completion 
				is probably already installed and operational.

		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
}
