import sqlite3

def get_user(username):
    conn = sqlite3.connect('database.db')
    cursor = conn.cursor()

    query = "SELECT * FROM users WHERE username = '" + username + "'"
    cursor.execute(query)

    user = cursor.fetchone()

    conn.close()

    return user

# Example usage
username = input("Enter a username: ")
user = get_user(username)

if user:
    print("User found:")
    print(user)
else:
    print("User not found.")
