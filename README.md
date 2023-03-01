# PriceBasket Application
The PriceBasket application is a command-line application that calculates the total price of a basket of items, including any discounts that apply. The application takes a list of item names as input and outputs the subtotal, discounts, and total price.

## Directory Organization
The PriceBasket solution is organized into two projects:
* **PriceBasket**: This project contains the application logic.
* **PriceBasket.Tests**: This project contains the unit tests for the application.

## Application Design
The PriceBasket application is structured in a way that follows the Model-View-Controller (MVC) pattern to guarrantee separation of concerns.

* **Models** - this folder contains the data models for the application, including the Product and Discount classes. 
* **Services** - this folder contains the classes responsible for data access, including ProductService and DiscountService.
* **Utils** - this folder contains utility classes for formatting, outputting data and handling business logic. 
* **Program.cs** - this file in the root directory contains the main entry point for the application and handles user input and output.

## Installation
To install the PriceBasket application, you need to have the .NET Core runtime installed on your system. You can download it from the official [.NET website](https://dotnet.microsoft.com/download).

Once you have .NET Core installed, you can then navigate to the `PriceBasket` directory and build the application using the following command:

```sh
dotnet build
```

This will build the application and create an executable file in the `bin/Debug/net6.0` directory.

## Usage
To use the PriceBasket application, run the following command:

```sh
dotnet run [item 1] [item 2] [item 3] ...
```

Replace `[item 1], [item 2], [item 3], ...` with the names of the items you want to include in the basket.

## Examples
Here are some example commands and their expected output:

#### Basket containing apples, milk, and bread
```sh
dotnet run Apples Milk Bread
```

Expected output:
```sh
Subtotal: £3.10
Apples 10% off: -10p
Total: £3.00
```

#### Basket containing only milk
```sh
dotnet run Milk
```

Expected output:
```sh
Subtotal: £1.30
(no offers available)
Total: £1.30
```

### Testing
#### Unit Tests
To run the unit tests for the PriceBasket application, navigate to the `PriceBasket.Tests` directory and run the following command:

```sh
dotnet test
```

This will execute all of the unit tests in the project and display the results in the console.

#### End-To-End Tests
The PriceBasket application comes with a suite of automated end-to-end tests that can be run using the following command:

```sh
.\end-to-end-test-1.bat
```
or
```sh
.\end-to-end-test-2.bat
```

This will run a series of tests to ensure that the application is functioning correctly. If all tests pass, the output should look like this:

```sh
SUBTOTAL TEST:      Passed
DISCOUNT TEST:      Passed
TOTAL TEST:         Passed
```

If any tests fail, the output will indicate which tests failed and why.
