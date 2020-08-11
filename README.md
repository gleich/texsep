<!-- DO NOT REMOVE - contributor_list:data:start:["Matt-Gleich"]:end -->

```txt
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

![build](https://github.com/Matt-Gleich/texsep/workflows/build/badge.svg)
![lint](https://github.com/Matt-Gleich/texsep/workflows/lint/badge.svg)
![test](https://github.com/Matt-Gleich/texsep/workflows/test/badge.svg)
![release](https://github.com/Matt-Gleich/texsep/workflows/release/badge.svg)

Separate your pdfs from your tex files

## Install

### MacOS/Linux

Simply run `brew tap Matt-Gleich/homebrew-taps` and then `brew install texsep`

## Structures

⚠️ All commands should be ran from the top of your LaTeX project ⚠️

### `texsep mirror`

See `texsep mirror --help` for more information

#### Initial Structure

Structure before the command runs

```txt
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

```txt
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

<!-- DO NOT REMOVE - contributor_list:start -->

## 👥 Contributors

- **[@Matt-Gleich](https://github.com/Matt-Gleich)**

<!-- DO NOT REMOVE - contributor_list:end -->
