<h1 align="center">Test Generator LLM</h1>


<div align="center">
  <!-- <img src="https://img.shields.io/badge/Assignment%20Project%20for%20-Keploy%20API%20Fellowship%20-ea580c?style=for-the-badge" alt="Showcase Project"> -->
  
  <!-- <br> -->
  
  <img src="https://img.shields.io/badge/License-MIT-ed8796.svg?style=for-the-badge" alt="MIT License">
</div>

<br>

> ### Use an LLM to generate test cases and execute them for validation.

We are using an LLM to generate unit tests for our code automatically. Once generated, the tests are run to check if they behave as expected and properly validate the functionality


## :computer: Tech Stack
- [**Go**](https://go.dev/) : Programming language
- [**Azure OpenAI/ChatGPT 4.1 Mini**](https://github.com/marketplace/models/azure-openai/gpt-4-1-mini): Utilized as the LLM backend to generate meaningful and contextual unit test cases for C++ code.
- [**CCP Project (orgChartApi)**](https://github.com/keploy/orgChartApi): The C++ project under test. The tool integrates with the CCP codebase to build the binary, generate test cases, and run them for validation.

<br>

## :book: How to Use / Run on Your Machine

- ### Prerequisites:
    - Install Go (version >= 1.23.3): https://golang.org/dl/
    - Setup the CCP Project locally (don't use docker to setup this project)
      - [Setup Instruction](https://github.com/keploy/orgChartApi?tab=readme-ov-file#%EF%B8%8F-manual-setup-without-docker)


> [!IMPORTANT]  
> I am using WSL (Ubuntu).

<br>

### :toolbox: Setup Project Locally:

- Clone the repository:
```sh
git clone https://github.com/saketv8/test-generator-llm.git
```

- Navigate to the project directory:
```sh
cd test-generator-llm
```

- Install dependencies:
```sh
go mod download
```

- Building the Application Binary:
```sh
go build -o test-generator-llm
```

- Starting the Application:
```sh
./test-generator-llm
```

- Check the terminal output
> :rocket: You're all set!

<br>

### :satellite: Terminal Output

<div align="center">
    <img src="./images/start_generation.png" alt="GET" style="width: 48%;">
    <img src="./images/build_100.png" alt="POST" style="width: 48%;">
</div>

<br>

<div align="center">
    <img src="./images/all_test_passed.png">
</div>

<br>

> [!TIP]  
> All tests passed (-_-)


### 📊 Viewing Generated Test

> Check the `./test/` folder


<br>

### 🐳 Program Design (overview)
<div align="center">
    <img src="./images/program_design.png">
</div>



<br>

### 🌸 Full terminal output
```sh

```

## :seedling: Todo / Future Improvements
- [x] Implemented the main functionality
- [ ] Code Refactoring

## :compass: About
This project was created as an assignment 5 for Keploy's API Fellowship Sessions.