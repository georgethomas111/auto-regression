# What is this?

Trying to get auto regression working with a simple library. 
Inspiration is coming from this youtube video.

https://www.youtube.com/watch?v=5-2C4eO4cPQ

# Requirements

## Overview
Given a timeseries graph and a particular time selected in the time period, 
draw the partial auto correlation function(pacf) for a point in the graph.

Steps:

* Write a function which accepts timestamps with values assosciated with the timestamp.
* If a point is returned in the timeseries and an amount to look backward, return the pacf in an array
For the index i and a given amount, return the partial auto correlation 
function for that.
