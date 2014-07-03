package main

type Config struct {
  Api struct {
    Url string
  }
  Credentials struct {
    Username string
    Password string
  }
}

type PapiApplicationResponse struct {
    Id                      string
    Name                    string
    Platform                string
    VoiceEnvironmentId      string
    MessagingEnvironmentId  string
    MessagingUrl            string
    VoiceUrl                string
    Environment             string
    UserId                  int
    Partition               string
}


type PapiUserResponse struct {
    Id   string
    Address     string
    Address2    string
    Username    string
    State       string
    Email       string
    FirstName   string
    LastName    string
    Notes       string
    Status      string
    JoinDate    string
    PasswordFailedAttempts int
}