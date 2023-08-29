#this test a vulnerable code
import hashlib

def hash_password(password):
    md5_hash = hashlib.md5(password.encode()).hexdigest()
    return md5_hash

# Example usage
user_password = input("Enter a password: ")
hashed_password = hash_password(user_password)
print("Hashed password:", hashed_password)
