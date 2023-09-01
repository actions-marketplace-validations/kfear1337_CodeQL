# CodeQL Security Workflow
[![GitHub](https://img.shields.io/github/license/kfear1337/CodeQL)](LICENSE)
[![GitHub last commit](https://img.shields.io/github/last-commit/kfear1337/CodeQL)](https://github.com/kfear1337/CodeQL/commits/main)
[![Python 3.x](https://img.shields.io/badge/Python-3.x-blue.svg)](https://www.python.org/downloads/)
[![Go](https://img.shields.io/badge/Go-1.x-blue.svg)](https://golang.org/dl/)
[![JavaScript](https://img.shields.io/badge/JavaScript-ES6-blue.svg)](https://developer.mozilla.org/en-US/docs/Web/JavaScript)
[![TypeScript](https://img.shields.io/badge/TypeScript-4.x-blue.svg)](https://www.typescriptlang.org/)
[![CodeQL-Security](https://github.com/kfear1337/CodeQL/actions/workflows/codeql.yml/badge.svg?event=schedule)](https://github.com/kfear1337/CodeQL/actions/workflows/codeql.yml)

This GitHub Actions workflow runs CodeQL security analysis on your repository. It analyzes the code for potential security vulnerabilities and generates SARIF reports for different programming languages.

## How it works?

#### Example for if someone pulls a request:

[Pull Request](https://github.com/kfear1337/CodeQL/pull/1)

[Code scanning results](https://github.com/kfear1337/CodeQL/pull/1/checks?check_run_id=16297913027)

#### It will show in your repository 

**example here:**

[Security Report](https://github.com/kfear1337/CodeQL/security/code-scanning?query=pr%3A1+tool%3ACodeQL+is%3Aopen)

![Example Image](https://i.imgur.com/ZygPoP8.png)

![Example Image](https://i.imgur.com/geaawB7.png)

### Read more:
[CodeQL : Code Scanning](https://docs.github.com/en/code-security/code-scanning/introduction-to-code-scanning/about-code-scanning-with-codeql)

## Workflow Description

The workflow is triggered on `push` and `pull_request` events on the `main` branch. It also runs daily at midnight (UTC) using a cron schedule.

The workflow consists of the following steps:

1. Checkout repository
2. Set up the required environment for the selected programming language (Python, JavaScript, or Go)
3. Install dependencies specific to the selected programming language
4. Detect the repository's primary language
5. Initialize CodeQL
6. Perform CodeQL security analysis
7. Force Upload SARIF reports for each programming language analyzed `(better than from Analyze default workflow)`
8. Download Result as `Artifacts` in the `Workflow` when each programming language analyzed complete [**NEW**]
9. The result, as `Artifacts`, is encrypted use `GPG Key` within a container of SARIF file. [**NEW**]

## Usage

To use this workflow in your repository, follow these steps:

1. Create a new file named `action.yml` inside the `.github/workflows` directory.
2. Copy the contents of the `action.yml` file from the repository you mentioned into the newly created `codeql.yml` file.
3. Create a new file named `codeql-config.yml` inside the `.github/workflows/codeql/` directory.
4. Copy the contents of your CodeQL configuration file into the newly created `.github/workflows/config/codeql-config.yml` file.
5. Customize the workflow file and the CodeQL configuration file as needed. You can adjust the timeout, permissions, and other settings according to your requirements.
6. Commit and push the changes to your repository.

The workflow will now be triggered on `push` and `pull_request` events on the `main` branch, as well as daily at midnight (UTC), based on the provided configuration.

Please make sure that the workflow file is located in the `.github/workflows` directory, and the CodeQL configuration file is located in the `.github/workflows/codeql` directory of your repository.

Feel free to modify the workflow file and the CodeQL configuration file to fit your specific needs and configurations.

The workflow will now be triggered on `push` and `pull_request` events on the `main` branch, as well as daily at midnight (UTC).

You can download the latest release of CodeQL from [here](https://github.com/kfear1337/CodeQL/releases).

## Additional Usage if you download from latest release [here](https://github.com/kfear1337/CodeQL/releases).

To use this workflow in your repository, follow these steps:

- Extract/Copy File Inside A Folder `workflows` into `.github/workflows`
- After extract/copy its should be `codeql.yml` and folder `codeql` for config `codeql-config.yml` Inside `.github/workflows`
- Customize the workflow file and the CodeQL configuration file as needed. You can adjust the timeout, permissions, and other settings according to your requirements.
- Commit and push the changes to your repository.

Please make sure that the workflow file is located in the `.github/workflows` directory, and the CodeQL configuration file is located in the `.github/workflows/codeql` directory of your repository.

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
