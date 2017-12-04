:io.format("~.2f~n", [3.1425])

IO.puts(System.get_env("PATH"))

IO.puts(Path.extname("hello.rs"))

IO.puts(System.cwd)

# https://github.com/devinus/poison

IO.inspect(System.cmd("ls", []))
