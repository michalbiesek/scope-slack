import os

from slack_sdk import WebClient
from slack_sdk.errors import SlackApiError

class Sender:

    def __init__(self) -> None:
        slack_bot_token = os.environ.get('SLACK_BOT_TOKEN')
        if not slack_bot_token:
            raise Exception("SLACK_BOT_TOKEN environment variable is not defined.")
        # Initialize the Slack WebClient with the token
        self.client = WebClient(token=slack_bot_token)

    def send_msg(self) -> None:
        try:
            response = self.client.chat_postMessage(channel='#test-app', text="Hello world!")
            assert response["message"]["text"] == "Hello world!"
        except SlackApiError as e:
            # You will get a SlackApiError if "ok" is False
            assert e.response["ok"] is False
            assert e.response["error"]  # str like 'invalid_auth', 'channel_not_found'
            print(f"Got an error: {e.response['error']}")
