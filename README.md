# Go tmux sessionizer

> [!WARNING]
> This project is under very active development, it works as expected in its current state, but major breaking changes may come later.

This is a CLI utility based on the `tmux-sessionizer` [script](https://github.com/ThePrimeagen/.dotfiles/blob/master/bin/.local/scripts/tmux-sessionizer) from ThePrimeagen. I've written this in Go, for fun and to learn Go.\
Running the script will display a list of directories, selecting one of the directories will create and switch to a new tmux session in that directory.
If a session already exists for the selected session, it will switch to the existing sesison instead of creating a new one.

## Config

It will look for `config.json` in `~/.config/tmux-sessionizer` in the following format.

```json
{
  "paths": ["~/", "~/Projects", "~/Documents", "~/Documents/notes"]
}
```

The sessionizer will display all the directories from the paths listed in the `paths` array.

## Usage

### Add the script to your `$PATH`

I've added the script to my `~/.local/bin` directory, and added that directory to my `$PATH`. That way I can easily add more scripts to my `$PATH` in the future.

You can do the same by adding this line to your `~/.zshrc` or `~/.bashrc` file (depending on which shell your using).

```bash
export PATH="$HOME/.local/bin:$PATH"
```

### Run from terminal

By adding this line to your `~/.zshrc` or `~/.bashrc` file (depending on which shell your using), you can run the script by pressing the `Ctrl`+`f` shortcut in your terminal.

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

- [`tmux`](https://github.com/tmux/tmux)
- [`fzf`](https://github.com/junegunn/fzf)

## Future improvements

- Better arguments with [cobra](https://github.com/spf13/cobra)
- Better config with [viper](https://github.com/spf13/viper)
- Initial sessions to be created
  - Configurable of course
  - When should they be created
    - Every time the script is run?
  - Need a way to reference the sessions
    - fzf will search directories and create a session based on directory name.
      - Pass a list of initial sessions to fzf, along with the paths supplied?
    - The initial sessions may depend on functionality, rather than directory location.
    - They may have custom names like "today" or "todo".
  - Example use cases
    - Daily note
    - Todo list
    - Dotfiles directory
- Plugin system?
