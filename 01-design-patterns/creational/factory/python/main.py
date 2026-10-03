import random
from abc import ABC, abstractmethod
from dataclasses import dataclass
from typing import Optional, Callable

@dataclass
class User:
    name: str
    user_id: str
    address: str | None = None

    # Instance Method
    # def greet(self) -> str:
    #     return f"Hello {self.name}"

    # @staticmethod
    # def add(a: int, b: int) -> int:
    #     return a + b

    # @classmethod
    # def from_dict(cls, data: dict):              # Behaving as Alternative Constructor - Since it gets cls not self
    #     return cls(data["name"],data["user_id"])

    # @property
    # def name(self):
    #     return self._name

    # @name.setter                                 # Property Setter only
    # def name(self, value):
    #     self._name = value

# from pydantic import BaseModel

# class User(BaseModel):
#     name: str
#     user_id: str
#     address: Optional[str] = None

class Notification(ABC):

    @abstractmethod
    def send(self, message, user):
        pass


class EmailNotification(Notification):

    def send(self, msg, user)->tuple[str, str | None]:
        if random.randint(0, 1) == 0:
            return "", "failed to send email"

        return (f"Email sent successfully to: {user.name}: {msg}",None)


class SmsNotification(Notification):

    def send(self, msg, user)->tuple[str, str | None]:
        if random.randint(0, 1) == 0:
            return "", "failed to send SMS"

        return (f"SMS sent successfully to: {user.name}: {msg}",None)


class PushNotification(Notification):

    def send(self, msg, user)->tuple[str, str | None]:
        if random.randint(0, 1) == 0:
            return "", "failed to send Push Notification"

        return (f"Push Notification sent successfully to: {user.name}: {msg}",None)


class NotificationFactory:

    def __init__(self):
        self.creators: dict[str, Callable[[], Notification]] = {}

    def register(self, notification_type, creator):
        self.creators[notification_type] = creator

    def create(self, notification_type):
        creator = self.creators.get(notification_type)

        if creator is None:
            return None

        return creator()


def main():
    user = User("Karan", "u1")
    user2 = User("Ramu", "u2")

    factory = NotificationFactory()

    factory.register("email",EmailNotification)
    factory.register("sms",SmsNotification)
    factory.register("push",lambda: PushNotification()) # Both ways work as Python class is callable

    email = factory.create("email")
    sms = factory.create("sms")
    push = factory.create("push")

    # Email
    res, err = email.send("Hello, World!", user)

    if err:
        print("Error:", err)
    else:
        print(res)

    # SMS
    res, err = sms.send("Hey there", user)

    if err:
        print("Error:", err)
    else:
        print(res)

    # Push
    res, err = push.send("Kya haal hain....", user2)

    if err:
        print("Error:", err)
    else:
        print(res)


if __name__ == "__main__":
    main()


# My Implementation

# from abc import ABC, abstractmethod
# from typing import TypedDict, Optional


# class User(TypedDict):
#     name: str
#     user_id: str
#     address: Optional[str]

# class Notification(ABC):
#     @abstractmethod
#     def Send(self, user: User, msg: str) -> None:
#         pass

# class EmailNotification(Notification):
#     def Send(self, user: User, msg: str) -> None:
#         """Send Email Notification"""
#         print(f"Sending email notification to: {user['name']}: {msg}")


# noti: Notification = EmailNotification()
# user: User = User(name="Karan",user_id="U1",address=None)

# noti.Send(user,"Hello, World!")