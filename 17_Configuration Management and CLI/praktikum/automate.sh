token=""
custom_date=`date +'%a %b %T WIB %Y'`
name_root="$1 at $custom_date"
architecture=`uname -a`
ping=`ping -c3 google.com`

mkdir "$name_root"
cd "$name_root"
mkdir "about_me"
cd "about_me"
mkdir "personal"
cd "personal"
echo "https://www.facebook.com/$2" >> facebook.txt
cd ..

mkdir "professional"
cd "professional"
echo "https://www.linkedin.com/in/$3" >> linkedin.txt
cd ..
cd ..
mkdir "my_friend"
cd "my_friend"
curl "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt >> list_of_my_friends.txt"
cd ..
mkdir "my_system_info"
cd "my_system_info"
echo "My username: $USER \nWith host: $architecture" >> about_this_laptop.txt
echo "$ping" >> internet_connection.txt