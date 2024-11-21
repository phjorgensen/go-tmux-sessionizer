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
        - In this case, I may not need to initialize the sessions until they are selected from the list.
    - The initial sessions may depend on functionality, rather than directory location.
    - They may have custom names like "today" or "todo".
  - Example use cases
    - Daily note
    - Todo list
    - Dotfiles directory
- Plugin system?

### Custom sessions that can be created

I'm noticing a problem with connecting the search directly with directory paths.

- What if you want to create two different sessions in the same directory?
- What if the session you want to create doesn't really have a direct connection to the directory itself?

My specific issue is when I'm using my note CLI tool, `nb`. It always works with notebooks I have defined in a specific directory, no matter what my current directory is, so the directory doesn't really matter.
It also supports local notebooks, which means that if I'm in a directory that is a notebook it will treat that notebook as the default and it will not let you change to another notebook. So I don't really want to be in the notebook directory either, as it block me from changing notebooks quickly.
So what I really need is some generic directory, like `$HOME`, with a custom session name so I can have multiple of them without stopping me from creating a session in `$HOME` if I would need that at some point.

A solution could be to add a way to specify custom sessions in the config file.

#### Config file

Define two custom sessions with the names "nb-work" and "nb-personal" that opens a session in `$HOME`.

```json
{
  "customSessions": {
    "nb-work": "~/",
    "nb-personal": "~/"
  }
}
```

#### Usage

This will open the session as if you had gone through the normal flow and selected this option from the `fzf`-list. This makes it really simple to bind to a key-bind in your shell.

```bash
go-tmux-sessionizer --custom "nb-work"
```
