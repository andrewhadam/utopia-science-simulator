# Utopia Science Simulator

The goal of this project was to create a quick script that could help me estimate how much total science a province might have by the end of an age given a couple of specific parameters

## Description

An in-depth paragraph about your project and overview of use.

## Getting Started

### Dependencies

* GoLang
* git

### Installing

* Install GoLang -  
https://go.dev/doc/install

* Checkout the code   
```
git clone https://github.com/andrewhadam/utopia-science-simulator.git 
```

* Build the code  

``` 
go build science.go 
```

### Executing program

* Update config.json  

The config.json controls the different parameters throughout the simulator.  You'll want to update the values to meet your needs.  

**ScientistGenerationRate**: This is the spawn rate of scientists  (Found: http://wiki.utopia-game.com/index.php?title=Science_Formulas)  
**WeeksInAge**: Total number of weeks in the age you want to simulate  
**RevelationUptime**: The percentage of time you expect Revelation to be up over the course of the age  (0 indicates you do not have Revelation)  
**FountainOfKnowledgeUptime**: The percentage of time you expect Fountain of Knowledge to be up over the course of the age (0 indicates you do not have FoK)  
**UniversityPercentage**: The average percentage of Universities you expect to have built when you have them    
**UniversityUptime**: The percentage of time you expect to have Universities built throughout the age  (0 indicates you expect to never have Universities built)  
**RacialScienceProductionModifier**: The racial modifier specific to your race pick (Currently does not function)  

* Run the Program
Either run the executable
```
// MacOSX / Linux 
./science

// Windows 
./science.exe
```

or 

```
go run science.go
```

## Version History

* 0.1
    * Initial Release