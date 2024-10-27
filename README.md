# Go TMUX sessionizer

> [!Notice]
> This project is under very active development, it works as expected in its current state, but major breaking changes may come later.

This is a CLI utility based on the `tmux-sessionizer` [script](https://github.com/ThePrimeagen/.dotfiles/blob/master/bin/.local/scripts/tmux-sessionizer) from ThePrimeagen. I've written this in Go, for fun and to learn Go.
Running the script will display a list of directories, selecting one of the directories will create and switch to a new tmux session in that directory.

## Config

It will look for a config file in `~/.config/tmux-sessionizer/config.json` in the following format.

```json
{
  "paths": ["~/", "~/Projects", "~/Documents", "~/Documents/notes"]
}
```

The sessionizer will display all the directories from the paths listed in the `paths` array.

## Usage

### Add the script to your $PATH

I've added the script to my `~/.local/bin` folder, and added that folder to my `$PATH`. That way I can easily add more scripts to my `$PATH` in the future.

### Run from terminal

By adding this line to your `~/.zshrc` or `~/.bashrc` file (depending on which shell your using.), you can run the script by pressing the `Ctrl`+`f` shortcut in your terminal.

```bash
bindkey -s ^f "go-tmux-sessionizer\n"
```

### Run from tmux

By adding this line to your `~/.config/tmux/tmux.conf`, you can run the script from an existing tmux session by pressing the `prefix`+`f` shortcut anywhere in the session.
The default tmux prefix is `Ctrl`+`b`, which means that pressing `Ctrl`+`b`, then `f`, will run the script.

```bash
bind-key -r f run-shell "tmux neww go-tmux-sessionizer"
```

## Dependencies

- [`tmux`](https://github.com/tmux/tmux/wiki)
- [`fzf`](https://github.com/junegunn/fzf)
