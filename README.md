# gitdl ↯

Download files and directories locally from a remote git repository

## Motivation

When updating my personal website, I needed a simple way to copy old blog posts
and their assets from an archived repository to the new working repository. I
didn't care about the version control revisions of these posts (most were never
edited after being created anyways), and I didn't want to have to clone the
entire archived repository and manually `mv` files/directories around.

Inspired by the Unix philosophy of building simple programs that do one thing
well like `cp` and `ftp`, I envisioned a simple CLI interface to copy files and
directories from a remote git repository to one's local filesystem. Since I was
learning Go at the time, the language seemed like a perfect candidate to utilize
for this project.

## Description

```
gitdl [OPTIONS]… REPO SOURCE [DEST_DIR]
```

- `REPO`: short GitHub path `<owner>/<repo>` (e.g. `cedricium/gitdl`)
- `SOURCE`: file(s) and/or director(y|ies) to download using a single string
  glob pattern
- `DEST_DIR`: local target destination directory where SOURCE args are copied
  to, default `.` or CWD

## Options

<!-- TODO -->

## Examples

The initial use-case that inspired development of `gitdl`:

```shell
gitdl cedricium/personal-site {posts,public/blog-assets} ~/git/website
```

## Development

> [!IMPORTANT]
>
> Requires Go v1.16+ since `go mod` is used to manage dependencies.

<!-- TODO: install/dev setup steps -->

### Implementation

Three ways to approach the problem of downloading remote repository contents
locally:

1. Using GitHub API to download desired files
2. Downloading archive of repository (tarball), unpacking and moving desired
   files
3. Utilizing
   [advanced `git` operations](https://github.blog/2020-01-17-bring-your-monorepo-down-to-size-with-sparse-checkout/)
   to selectively fetch desired files

<!-- TODO: explain pros/cons of each approach, why landed on 3 -->

## License

[MIT License](LICENSE.md)
