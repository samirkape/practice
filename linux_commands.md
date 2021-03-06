```bash
# how to compare which version of the linux program is installed
yq_version=`yq --version | head -n1 | cut -d" " -f3`

version_greater_equal()
{
    printf '%s\n%s\n' "$2" "$1" | sort --check=quiet --version-sort
}

version_greater_equal "${yq_version}" 2.2.1 || { echo "found yq ${yq_version} need 2.2.1"; exit -1; }
```

```bash
#how to find and replace in vim editor
:[range]s/{pattern}/{replace}/[flags] [count]
e.g
:%s/pattern/replace/g

# :s = substitute
# %  = range delimiter. If you want to search and replace the pattern in the entire file.
# g  = replace all the occurances in current line. If not specified, only first ocurrance will be replaced  
```

```bash
$ echo "Samir Nitin Kape" | cut -d" " -f3
$ Kape

$ echo "Samir Nitin Kape" | cut -d" " -f1
$ Samir

-d" " # Space is a Delimiter 
```

```bash
1. whereis -b yq
# finds all programs with name yq 
# export WHEREIS_DEBUG=all will print debug output 
2. type -a yq
```

```bash
info command 
# produces better command information than man 
e.g info grep
```

```bash
cut -f2 -d, csv_file.csv
#-f = column number 
#-d = delimiter
#cuts the specified column out of csv
```

 ```bash
 mkdir -p gp/parent/child  
 # -p = while creating child, 
 # creates gp and parent if not already present
 ```

 ```bash
#check the return type of any command executed
e.g 
mv x y 
result=$?   
if [ $result -ne 0 ]    # check if move is successful
then
    echo " "
    echo "error occured!"
    exit $result
fi
 ```

 ```bash
find . -type f -name "*.yaml" -exec echo {} \; -exec yq r {} version \; -exec echo "" \;
 ```

 ```bash
date +"%c"  #Current date-time-year
 ```

 ```bash
if [ "$#" -eq 0 ] #check number of input arguments
 ```

 ```bash
var=1 && for i in *.wav; do mv $i "noise_ahh_${var}.wav"; var=$((var+1)); done
  #rename files 1..n
 ```

 ```bash
X="filename.txt"
${X%.*}  -- filename 
 ```

 ```bash
sed -i 's/search/replace/g' file.txt
  #-i - inplace
 ```

 ```bash
sed -i '/delete-line-by-text-token/d' file.txt 
#No substitute (s/) at start
 ```

 ```bash
sed -i '/*----*/d' file.txt
  #asterisk = wildcard 
 ```

```bash
#wildcards
Wild-cards are handled completely by the shell before the associated
program even runs

* -- Zero or more consecutive characters

? -- Any single character

[set] -- Any single character in the given set, most commonly a 
	sequence of characters, like [aeiouAEIOU] for all vowels, or 
	a range with a dash, like [A-Z] for all capital letters

[^set] -- Any single character not in the given set, such as [^0-9] to mean any nondigit
```

 ```bash
head -n 1 filename 
 ```

 ```bash
scp -i pem_file.pem user@ip_addr:/path/to/the/file  out_dir 
 ```

 ```bash
scp username@ip_addr:/path/to/the/file out_dir 
 ```

 ```bash
var="bar"
echo "foo_${var}
# "foo_$var" does not work while concatenating  
 ```

 ```bash
for file in `cat $csv_filename`; do echo $file; done
#this will iterate over filenames provided in the csv_filename
 ```

```bash
export PS1="\[\e[32m\]\u@\h \[\e[34m\]\W \[\e[32m\]$ "
```

### VSCode Shortcuts
```bash
  1. Ctrl + Shift + \   # Jump to matching bracket
  2. Shift + Alt + A    # Block comment
  3. Alt + Z            # Word wrap
  4. CMD + P            # Show files
  5. F8                 # Goto next error or warning 
  6. Alt + Shift + F    # Format Document
  7. Ctrl + `           # Show terminal
```

### macOS Shortcuts
```bash
CMD + , # Opens app settings

```

<details><summary>Terminal Shortcuts List</summary>

* `Ctrl+a` Move cursor to start of line
* `Ctrl+e` Move cursor to end of line
* `Ctrl+b` Move back one character
* `Alt+b` Move back one word
* `Ctrl+f` Move forward one character
* `Alt+f` Move forward one word
* `Ctrl+d` Delete current character
* `Ctrl+w` Cut the last word
* `Ctrl+k` Cut everything after the cursor
* `Alt+d` Cut word after the cursor
* `Alt+w` Cut word before the cursor
* `Ctrl+y` Paste the last deleted command
* `Ctrl+_` Undo
* `Ctrl+u` Cut everything before the cursor
* `Ctrl+xx` Toggle between first and current position
* `Ctrl+l` Clear the terminal
* `Ctrl+c` Cancel the command
* `Ctrl+r` Search command in history - type the search term
* `Ctrl+j` End the search at current history entry
* `Ctrl+g` Cancel the search and restore original line
* `Ctrl+n` Next command from the History
* `Ctrl+p` previous command from the History

</details>