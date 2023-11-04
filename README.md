# grassflow (deprecated)

**Use shell script instead**

Run these two script in order

```shell
#!/bin/bash

# Clear the output.txt file
> output.txt

# Get all files in the current directory
find . -type f | while read -r file; do
  # Extract the base name of the file to exclude hidden files (starting with .)
  base_name=$(basename "$file")
  # Check if the file should be ignored by .gitignore, and ensure it has a non-hidden extension
  if ! git check-ignore -q "$file" ; then
    # Get the last modified time in ISO 8601 format
    mod_time=$(stat -f "%Sm" -t "%Y-%m-%d %H:%M:%S" "$file")
    # Output the last modified time and file path to output.txt
    echo "$mod_time $file" >> output.txt
  fi
done

# Sort the output.txt by date
sort -o output.txt output.txt
```

```shell
#!/bin/bash

while IFS= read -r line; do
  # Extract the date and filepath
  commit_date_str=$(echo "$line" | awk '{print $1 " " $2}')
  file_path=$(echo "$line" | cut -d' ' -f3-)
  
  # Convert the date to the expected format for Git
  GIT_COMMITTER_DATE=$(date -j -f "%Y-%m-%d %H:%M:%S" "$commit_date_str" +"%Y-%m-%d %H:%M:%S")
  export GIT_COMMITTER_DATE

  # Add the file to the staging area
  git add "$file_path"

  # Commit the file with a message containing the date
  git commit --date="$GIT_COMMITTER_DATE" -m "Commits at $GIT_COMMITTER_DATE"
done < output.txt
```
