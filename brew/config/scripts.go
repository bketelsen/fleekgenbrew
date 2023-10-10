package config

var (
	scripts map[string][]string
)

func init() {
	scripts = make(map[string][]string)

	scripts["install-astro"] = []string{
		"echo cloning astronvim repo to ~/.config/nvim",
		"git clone --depth 1 https://github.com/AstroNvim/AstroNvim ~/.config/nvim",
	}
	scripts["npm-global"] = []string{
		"grep -qF '.npm-packages' ~/.npmrc || echo 'prefix=~/.npm-packages' >> ~/.npmrc",
		"grep -qF '.npm-packages' ~/.zshrc || echo 'export PATH=$PATH:~/.npm-packages/bin' >> ~/.zshrc",
		"grep -qF '.npm-packages' ~/.bashrc || echo 'export PATH=$PATH:~/.npm-packages/bin' >> ~/.bashrc",
		"mkdir -p ~/.npm-packages",
	}
	scripts["install-starship-prompt"] = []string{
		"if [ ! -f $HOME/.config/starship.toml ]",
		"then",
		"curl -fsSL https://devbox.getfleek.dev/config/starship/starship.toml > $HOME/.config/starship.toml",
		"fi",
	}

}
