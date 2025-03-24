import google.auth
from google.auth.transport.requests import Request
from google.oauth2 import service_account

key_path = 'C:\\Users\\Shivansh sharma\\Downloads\\webpush-51b10-firebase-adminsdk-fbsvc-8316c67cd8.json'

credentials = service_account.Credentials.from_service_account_file(
    key_path,
    scopes=["https://www.googleapis.com/auth/firebase.messaging"]   
)

credentials.refresh(Request())

print(f'Access Token: {credentials.token}')