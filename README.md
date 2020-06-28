```
                    ___           ___            ___           ___           ___
                   /\__\         /|  |          /\__\         /\__\         /\  \
      ___         /:/ _/_       |:|  |         /:/ _/_       /:/ _/_       /::\  \
     /\__\       /:/ /\__\      |:|  |        /:/ /\  \     /:/ /\__\     /:/\:\__\
    /:/  /      /:/ /:/ _/_   __|:|__|       /:/ /::\  \   /:/ /:/ _/_   /:/ /:/  /
   /:/__/      /:/_/:/ /\__\ /::::\__\_____ /:/_/:/\:\__\ /:/_/:/ /\__\ /:/_/:/  /
  /::\  \      \:\/:/ /:/  / ~~~~\::::/___/ \:\/:/ /:/  / \:\/:/ /:/  / \:\/:/  /
 /:/\:\  \      \::/_/:/  /      |:|~~|      \::/ /:/  /   \::/_/:/  /   \::/__/
 \/__\:\  \      \:\/:/  /       |:|  |       \/_/:/  /     \:\/:/  /     \:\  \
      \:\__\      \::/  /        |:|__|         /:/  /       \::/  /       \:\__\
       \/__/       \/__/         |/__/          \/__/         \/__/         \/__/
```

# texsep

![Go Build](https://github.com/Matt-Gleich/texsep/workflows/Go%20Build/badge.svg) ![Go Format](https://github.com/Matt-Gleich/texsep/workflows/Go%20Format/badge.svg) ![goreleaser format](https://github.com/Matt-Gleich/texsep/workflows/goreleaser%20format/badge.svg) ![goreleaser](https://github.com/Matt-Gleich/texsep/workflows/goreleaser/badge.svg)

Separate your pdfs from your tex files

## Install

### MacOS/Linux

`brew install Matt-Gleich/homebrew-taps/texsep`

## Structures

⚠️ All commands should be ran from the top of your LaTeX project ⚠️

### `texsep mirror`

See `texsep mirror --help` for more information

#### Initial Structure

Structure before the command runs

```
├── Physics/
│   ├── string-theory.tex
│   ├── string-theory.pdf
│   ├── projectile-motion.tex
│   └── projectile-motion.pdf
│
├── Math/
│   ├── combinatorics.tex
│   ├── combinatorics.pdf
│   ├── ray-tracing.tex
│   └── ray-tracing.pdf
│
├── Computational-Theory/
│   ├── parallel-computing.tex
│   ├── parallel-computing.pdf
│   ├── microprocessors.tex
│   └── microprocessors.pdf
│
├── random.tex
└── random.pdf
```

#### Post Structure

Structure after the command runs

```
├── TeX/
│   ├── random.tex
│   ├── Physics/
│   │   ├── string-theory.tex
│   │   └── projectile-motion.tex
│   ├── Math/
│   │   ├── combinatorics.tex
│   │   └── ray-tracing.tex
│   └── Computational-Theory/
│       ├── parallel-computing.tex
│       └── microprocessors.tex
│
└── PDF/
    ├── random.pdf
    ├── Physics/
    │   ├── string-theory.pdf
    │   └── projectile-motion.pdf
    ├── Math/
    │   ├── combinatorics.pdf
    │   └── ray-tracing.pdf
    └── Computational-Theory/
        ├── parallel-computing.pdf
        └── microprocessors.pdf
```

<!-- ### `texsep folders`

See `texsep folders --help` for more information

#### Initial Structure

Structure before the command runs

```
├── Physics/
│   ├── string-theory.tex
│   ├── string-theory.pdf
│   ├── projectile-motion.tex
│   └── projectile-motion.pdf
│
├── Math/
│   ├── combinatorics.tex
│   ├── combinatorics.pdf
│   ├── ray-tracing.tex
│   └── ray-tracing.pdf
│
├── Computational-Theory/
│   ├── parallel-computing.tex
│   ├── parallel-computing.pdf
│   ├── microprocessors.tex
│   └── microprocessors.pdf
│
├── random.tex
└── random.pdf
```

#### Post Structure

Structure after the command runs

```
├── Physics/
│   ├── TeX/
│   │   ├── string-theory.tex
│   │   └── projectile-motion.tex
│   └── PDF/
│       ├── string-theory.pdf
│       └── projectile-motion.pdf
│
├── Math/
│   ├── TeX/
│   │   ├── combinatorics.tex
│   │   └── ray-tracing.tex
│   └── PDF/
│       ├── combinatorics.pdf
│       └── ray-tracing.pdf
│
├── Computational-Theory/
│   ├── TeX/
│   │   ├── parallel-computing.tex
│   │   └── parallel-computing.pdf
│   └── PDF/
│       ├── microprocessors.tex
│       └── microprocessors.pdf
│
├── TeX/
│   └── random.tex
│
└── PDF/
    └── random.pdf
``` -->
