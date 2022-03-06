# Mimic

The open-source Golang HTML templating engine

## About

The Mimic templating engine is based on
the [Django templating engine](https://github.com/django/django/blob/a46bc327e70f81b66800780edf3830f6137a89e3/django/template/base.py)
, meaning its parser creates an interpretation model for the renderer to consume.

## Process Overview

1. The Lexer tokenizes the template string, which is passed from the caller, to allow for a decoupled template storage
   mechanism.
2. The Parser parses the tokens into an abstract Node tree, which is then passed to the Renderer for rendering.
3. The Renderer renders the abstract Node tree into a string.

- Each Node has a render function, which outputs the string to be concatenated into the final render.

## Features

| Feature              | Description                                           | Implemented | Implemented In |
|----------------------|-------------------------------------------------------|-------------|----------------|
| Template inheritance | Templates can inherit master templates / layouts      | :x:         |                |
| Template imports     | Templates can import other templates                  | :x:         |                |
| Model binding        | Templates can bind models to the template context     | :x:         |                |
| Variable tokens      | Variables can be used in the template string          | :x:         |                |
| Conditional blocks   | Conditional blocks can be used in the template string | :x:         |                |
| Loop blocks          | Loop blocks can be used in the template string        | :x:         |                |
| Function blocks      | Function blocks can be used in the template string    | :x:         |                |
| Filter blocks        | Filter blocks can be used in the template string      | :x:         |                |

## Roadmap

| Version | Features                                                 | Released | Date of Release |
|---------|----------------------------------------------------------|----------|-----------------|
| 0.1.0   | Variable tokens<br />Conditional blocks<br />Loop blocks | :x:      |                 |
| 0.2.0   | Function blocks<br />Filter blocks                       | :x:      |                 |
| 0.3.0   | Template inheritance                                     | :x:      |                 |
| 0.4.0   | Template imports                                         | :x:      |                 |
| 0.5.0   | Model binding                                            | :x:      |                 |

## Contributing

Everyone is welcome to contribute to this project! From fixing bugs, adding features, or improving the documentation,
all PRs are considered.

Commit style:

```
feat(<feature name>): <feature description> OR
bug(<bug name>): <bug description> OR
maintenance(<maintenance name>): <maintenance description>

<commit overview>

closes #123 OR
updates #123

<list of updates in commit>
```

For more information on how to style the list of updates in the commit, please
consult [this guide](https://gist.github.com/parmentf/035de27d6ed1dce0b36a).

## Acknowledgements

[1] - [Django](https://www.djangoproject.com/), the major inspiration for this project.

[2] - [A Template Engine by Ned Batchelder](https://aosabook.org/en/500L/a-template-engine.html), whose article
explained what a templating engine does, and the different ways of doing it, in a digestable format.

## License

[Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0)