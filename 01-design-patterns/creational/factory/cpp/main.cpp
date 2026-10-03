#include <iostream>
#include <string>
#include <unordered_map>
#include <functional>

using namespace std;

class User {
public:
    string name;
    string userId;

    User(string name, string userId) {
        this->name = name;
        this->userId = userId;
    }
};

class Notification {
public:
    virtual void send(string msg, User user) = 0;
};

class EmailNotification : public Notification {
public:
    void send(string msg, User user) override {
        cout << "Email sent to " << user.name << ": " << msg << endl;
    }
};

class SmsNotification : public Notification {
public:
    void send(string msg, User user) override {
        cout << "SMS sent to " << user.name << ": " << msg << endl;
    }
};

class PushNotification : public Notification {
public:
    void send(string msg, User user) override {
        cout << "Push sent to " << user.name << ": " << msg << endl;
    }
};

class NotificationFactory {
private:
    unordered_map<string, function<Notification*()>> creators;

public:
    void registerCreator(string type, function<Notification*()> creator) {
        creators[type] = creator;
    }

    Notification* create(string type) {
        return creators[type]();
    }
};

int main() {
    User user("Karan", "U1");

    NotificationFactory factory;

    factory.registerCreator("email", []() {
        return new EmailNotification();
    });

    factory.registerCreator("sms", []() {
        return new SmsNotification();
    });

    factory.registerCreator("push", []() {
        return new PushNotification();
    });

    Notification* email = factory.create("email");
    Notification* sms = factory.create("sms");
    Notification* push = factory.create("push");

    email->send("Hello", user);
    sms->send("Hello", user);
    push->send("Hello", user);

    delete email;
    delete sms;
    delete push;
}