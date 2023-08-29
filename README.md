# CodeQL Security Workflow
[![GitHub](https://img.shields.io/github/license/kfear1337/CodeQL)](LICENSE)
[![GitHub last commit](https://img.shields.io/github/last-commit/kfear1337/CodeQL)](https://github.com/kfear1337/CodeQL/commits/main)
[![Python 3.x](https://img.shields.io/badge/Python-3.x-blue.svg)](https://www.python.org/downloads/)
[![Go](https://img.shields.io/badge/Go-1.x-blue.svg)](https://golang.org/dl/)
[![JavaScript](https://img.shields.io/badge/JavaScript-ES6-blue.svg)](https://developer.mozilla.org/en-US/docs/Web/JavaScript)
[![TypeScript](https://img.shields.io/badge/TypeScript-4.x-blue.svg)](https://www.typescriptlang.org/)


This GitHub Actions workflow runs CodeQL security analysis on your repository. It analyzes the code for potential security vulnerabilities and generates SARIF reports for different programming languages.

## How it works?

Example for if an someone pull request:

[Pull Request](https://github.com/kfear1337/CodeQL/pull/1)

It will show here:
[Security Report](https://github.com/kfear1337/CodeQL/security/code-scanning?query=pr%3A1+tool%3ACodeQL+is%3Aopen)

Read more:
[CodeQL Advance Security](https://docs.github.com/en/code-security/code-scanning/introduction-to-code-scanning/about-code-scanning)

## Workflow Description

The workflow is triggered on `push` and `pull_request` events on the `main` branch. It also runs daily at midnight (UTC) using a cron schedule.

The workflow consists of the following steps:

1. Checkout repository
2. Set up the required environment for the selected programming language (Python, JavaScript, or Go)
3. Install dependencies specific to the selected programming language
4. Detect the repository's primary language
5. Initialize CodeQL
6. Perform CodeQL security analysis
7. Upload SARIF reports for each programming language analyzed

## Usage

To use this workflow in your repository, follow these steps:

1. Create a new file named `.github/workflows/codeql.yml` in your repository.
2. Copy the contents of the `codeql.yml` file from the repository you mentioned into the newly created `codeql.yml` file.
3. Create a new directory named `.github/codeql` in your repository.
4. Create a new file named `codeql-config.yml` inside the `.github/codeql` directory.
5. Copy the contents of your CodeQL configuration file into the newly created `.github/codeql/codeql-config.yml` file.
6. Customize the workflow file and the CodeQL configuration file as needed. You can adjust the timeout, permissions, and other settings according to your requirements.
7. Commit and push the changes to your repository.

The workflow will now be triggered on `push` and `pull_request` events on the `main` branch, as well as daily at midnight (UTC), based on the provided configuration.

Please make sure that the workflow file is located in the `.github/workflows` directory, and the CodeQL configuration file is located in the `.github/codeql` directory of your repository.

Feel free to modify the workflow file and the CodeQL configuration file to fit your specific needs and configurations.

The workflow will now be triggered on `push` and `pull_request` events on the `main` branch, as well as daily at midnight (UTC).

## Configuration

The workflow can be configured to analyze different programming languages by modifying the `matrix.language` field in the workflow file. The currently supported languages are Python, JavaScript, and Go.

To add support for additional programming languages:

1. Edit the `matrix.language` field in the workflow file to include the new language.
2. Add steps to set up the required environment and install dependencies for the new language.
3. Modify the CodeQL analysis step to include the new language.
4. Update the SARIF report names in the "Results" section to include the new language.

You can customize the workflow further by adjusting the timeout, permissions, and other settings as needed.

## Results

The workflow generates SARIF reports for each programming language analyzed. The reports are uploaded as artifacts and can be accessed under the "Actions" tab in your repository.

- JavaScript SARIF: Contains the results of the JavaScript analysis.
- Python SARIF: Contains the results of the Python analysis.
- Go SARIF: Contains the results of the Go analysis.
- [Add SARIF report names for additional languages here]

## Permissions

The workflow requires the following permissions:

- `actions: read`
- `contents: read`
- `pull-requests: read`
- `deployments: read`
- `security-events: write`

Make sure the necessary permissions are granted to the workflow for it to run successfully.

## Roadmap

Here's a roadmap for adding support for other programming languages:

1. Identify the programming language you want to add support for.
2. Research the CodeQL setup and analysis requirements for the language.
3. Modify the workflow file to include the new language in the `matrix.language` field.
4. Add steps to set up the required environment and install dependencies for the new language.
5. Modify the CodeQL analysis step to include the new language.
6. Update the SARIF report names in the "Results" section to include the new language.

To add support for a new programming language:

1. Identify the programming language you want to add support for. For example, let's say you want to add support for Ruby.
2. Research the CodeQL setup and analysis requirements for Ruby. Check the CodeQL documentation for any language-specific instructions.
3. Modify the workflow file (`codeql.yml`) to include the new language in the `matrix.language` field. Add `'ruby'` to the list of supported languages.
4. Add steps to set up the required environment and install dependencies for Ruby. For example, you might need to install Ruby and any necessary gems.
5. Modify the CodeQL analysis step to include the new language. Update the `languages` field in the `Initialize CodeQL` step to include `'ruby'`.
6. Update the SARIF report names in the "Results" section to include the new language. For example, you can add a line like `- Ruby SARIF: Contains the results of the Ruby analysis.`.

Please note that the current version of this repository only supports Python, JavaScript, TypeScript, and Go. If you want to add support for additional languages, you'll need to follow the steps mentioned above and adjust the workflow file accordingly.


## License

This workflow is licensed under the [MIT License](LICENSE).

## My Gist

You can find my Gist on GitHub at the following link:
[gist.github.com/kfear1337/4f1c754aba3dd66a8c463b26695eba56](https://gist.github.com/kfear1337/4f1c754aba3dd66a8c463b26695eba56)
