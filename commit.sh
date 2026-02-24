echo "ENTER COMMIT MESSAGE"
read msg

git add .
git commit -m "$msg"
git push