Project Hangman

Poré Axel, Capo Arthur, Noniashvili Zezva

Usage: ./<program_name> [file_word_path]

First, the program will collect a word in a .txt file.

Seconde, the program create an hiddenword from word who reveal n letter, n worth the length of word divide by 2 - 1.

Third, the program will demand a letter to the user and compare the recieve letter with the letter in the word, if a comparison exist 

the program display the letter(s) in the hiddenword and return a true.

Else he will return a false.

Forth, if the program recieve a true, he will enter in win and display the hiddenword, then he will returns to the Third.

Else if he recieve a false, the program will display "not present in the word", the remaining attemps and José. Next the program return

in Third.

Last, the program display a win or game-over message.

Example:

Good Luck, you have 10 attempts.

_ _ _ _ O

Choose: E

_ E _ _ O

Choose: A

Not present in the word, 9 attempts remaining
         
         
         
         
         
         
=========


Choose: L

_ E L L O 

Choose: B

Not present in the word, 8 attempts remaining
         
      |  
      |  
      |  
      |  
      |  

=========

Choose: Z

Not present in the word, 7 attempts remaining

  +---+  

      |  
      |  
      |  
      |  
      |  
      
=========

Choose: H

H E L L O

Congrats !