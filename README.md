
# Sleep Tracker Application

This project is aimed to track users sleep habits and suggest activities to the user to improve sleep quality and health. In this project, users only have to input date, sleeping and wake up time on a specific date, and the program is responsible for storing and calculating both sleep duration and sleep quality automatically. The calculation will be used for sleep improvement suggestions.
This project is made 100% using Go Programming Language and can be executed using terminal or Integrated Development Enviroment (IDE).


## Demo

https://drive.google.com/drive/folders/1WtER52P04Wml6C3MZ92yxYRxaab2rUWa?usp=sharing


## Features & Technical Aspects

- Add Records

    Users can add sleep records with the date,  sleep time, wake-up time, and the system calculates the sleep duration and quality. It will stored in an array of struct sleepRecord.
- Modify Records

    Users can modify an existing record by specifying the date they want to edit. The program will then find the matching record via sequential search. After that, the users can change the date, sleep time, or wake-up time and the program will adjust the sleep duration and quality based on new input.
- Sleep Duration Calculation

    The system calculates sleep duration based on the difference between sleep time and wake-up time, crosses over 24-hour boundary.
- Sleep Quality Calculation

    The application calculates sleep quality based on the following factors:

        - Sleep duration, where the optimal sleep duration is between 7-9 hours.
        - Sleep time, where the optimal time to sleep is before 23:00.
        - Suboptimal and oversleeping conditions are handled with specific quality ratings.
- Delete Records

    Users can delete a sleep record by specifying the date and the program will find the matching record via sequential search. The program will delete the record by shifting subsequent records to fill the gap.
- Search Records by Date

    Users can search for a sleep record by entering a specific date. The search can be performed using binary search after the records are sorted by date.
- Display Records

    The system displays the records with the following details:
    
        - Date
        - Sleep time
        - Wake up time
        - Calculated sleep duration
        - Calculated sleep quality rating
- Sorted Records

    Users can view all sleep records with options to sort in ascending order using insertion sort, or in descending order using selection sort by either:
    
        - Date
        - Sleep Duration
- Weekly Summary

    Users can view a summary of the last 7 recorded days of sleep data, which includes:

        - Date 
        - Sleep time 
        - Wake-up time
        - Calculated sleep duration
        - Calculated sleep quality rating
        - Average sleep duration (in hours)
        - Average sleep quality rating
        - Personalized suggestions for improving sleep quality
- Suggestion for Sleep Improvement

    Based on the weekly summary, average sleep time and quality, the system recommends actions to improve sleep quality (e.g., sleep before a certain hour, maintain or adjust sleep duration).


## Target Scenario Example

Input the following 8 sleep records into the program.

Date format: yy/mm/dd 

Time format: h m (24-hour clock)

    1. 25/05/26
        - Sleep time: 23 0
        - Wake-up time: 5 0

    2. 25/05/27
        - Sleep time: 22 30
        - Wake-up time: 8 10

    3. 25/05/28
        - Sleep time: 21 15
        - Wake-up time: 8 10

    4. 25/05/29
        - Sleep time: 22 0
        - Wake-up time: 7 30

    5. 25/05/30
        - Sleep time: 23 30
        - Wake-up time: 8 15

    6. 25/05/31
        - Sleep time: 0 15
        - Wake-up time: 7 45

    7. 25/06/01
        - Sleep time: 23 10
        - Wake-up time: 6 15

    8. 25/06/02
        - Sleep time: 21 0
        - Wake-up time: 5 20

Then, view the records sorted by sleep duration in descending order. You will be able to view the date that has the shortest and longest sleep duration. Next, adjust the wake-up time of the longest sleeping duration to 6:50. Lastly, view the weekly summary and suggestions to improve your sleep quality.

## Step-by-Step Instructions
This is the main menu of the program:

    --- Sleep Tracker ---
    1. Add Record
    2. Modify Record
    3. Delete Record
    4. Search Record by Date
    5. View Records
    6. View Weekly Summary
    0. Exit

In order to execute the above scenarion, you will have to follow the given instructions.

    Add Records
    1. Choose option 1 to Add Record.
    2. When prompted, enter 8 to indicate how many records you want to add.
    3. Enter the 8 records as shown in the list above.

    View Sorted Records
    4. After input is complete, the system returns to the home page.
    5. Choose option 5 to View Records.
    6. Choose option 2 to sort records by Sleep Duration.
    7. Choose option 2 again to sort the duration in Descending Order.
    8. The record with the longest sleep duration will be at the top, and the shortest at the bottom.
    9. Enter 0 to exit the View Records page.

    Modify Wake-Up Time
    10. Back on the home page, choose option 2 to Modify Record
    11. From the sorted list, the record with the longest sleep duration has the date 25/05/28.
    12. Enter 25/05/28 as the date to modify.
    13. Choose option 3 to edit Wake-Up Time.
    14. Input the new wake-up time: 6 50
    15. Enter 0 to return to the Modify Record menu.
    16. Enter x to return to the Home Page.

    View Weekly Summary & Suggestions
    17. Choose option 6 to View Weekly Summary.
    18. When asked if you'd like to view suggestions, enter yes
    19. The system will display a summary and tips to improve your sleep quality.
    20. After viewing, you'll be returned to the Home Page.




## Lessons Learned

The most challenging features to develop was “Sorting by Date”, because the date is stored as a string rather than an integer. This challenge took a considerable amount of time to solve, which slowed down the overall development of the application.

There were two possible solutions for this issue:

    1. Using the time package, which provides functionality for for measuring and displaying date and time.
    2. Comparing string values manually by using the yy/mm/dd date format. This format allows the system to compare the year, then the month, and finally the day from left-to-right, enabling accurate date sorting without converting to another data type.

Since the "Sorting by Date" is a challenging feature to develop, it affects the development time of searching records via binary search. However after implementing the string values comparison as a solution for "Sorting by Date", this feature was quickly solved and coded.

Moreover, one of the key challenges in implementing feature modifications is accurately determining which specific information to update. This complexity arises when users choose to modify only certain data points, such as sleep quality, rather than the entire dataset. Addressing this challenge requires implementing targeted data modification logic that isolates and updates the specified data while maintaining data integrity throughout the system. 

