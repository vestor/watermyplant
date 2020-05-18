# Water My Plant
Given the abject aversion to take responsibilities, 
this code exists to water my flatmate's plant. 
This is the backend service which operates solely from the raspberry-pi which is giving the plant company.

Written in go, using pi-blaster it operates a stepper motor to water the plant.

## Why?
As mentioned before, this exists to water my flatmate's plant.
This is also an exercise to: 
1. Learning GoLang for starters.
2. Interfacing with Raspberry-Pi using pi-blaster
3. Discipline for personal projects

## How?
Later

### TODO
 -[x] Implement the authentication
    -[x] Use JWT 
    -[x] Implement secure APIs
 -[ ] Add license
 -[ ] Add "How" section in README 
 -[ ] Move to REDIS as the persistence store
 -[ ] Use time-series for water logs 
 -[ ] Enable live-streaming of plant
 -[ ] Stress test the system