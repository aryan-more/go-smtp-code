package smtp_codes

var SystemStatus = SmtpCode{Code: 211, EnhancedCodes: NoEnhancedCodes}

var Help = SmtpCode{Code: 214, EnhancedCodes: NoEnhancedCodes}

var ServiceReady = SmtpCode{Code: 220, EnhancedCodes: NoEnhancedCodes}

var ServiceClosingTransmissionChannel = SmtpCode{Code: 221, EnhancedCodes: NoEnhancedCodes}

var Goodbye = SmtpCode{Code: 221, EnhancedCodes: [3]int16{}}

var AuthenticationSucceeded = SmtpCode{Code: 235, EnhancedCodes: [3]int16{}}

var Quit = SmtpCode{Code: 240, EnhancedCodes: NoEnhancedCodes}

var RequestedMailActionCompleted = SmtpCode{Code: 250, EnhancedCodes: NoEnhancedCodes}

var UserNotLocalButWillForward = SmtpCode{Code: 251, EnhancedCodes: NoEnhancedCodes}

var CannotVerifyTheUserButTryToDeliver = SmtpCode{Code: 252, EnhancedCodes: NoEnhancedCodes}

var ServerChallenge = SmtpCode{Code: 334, EnhancedCodes: NoEnhancedCodes}

var StartMailInput = SmtpCode{Code: 354, EnhancedCodes: NoEnhancedCodes}

var ServiceNotAvailable = SmtpCode{Code: 421, EnhancedCodes: NoEnhancedCodes}

var PasswordRequired = SmtpCode{Code: 432, EnhancedCodes: [3]int16{}}

var TemporaryMailboxUnavailable = SmtpCode{Code: 450, EnhancedCodes: NoEnhancedCodes}

var RequestedActionAborted = SmtpCode{Code: 451, EnhancedCodes: NoEnhancedCodes}

var ImapServerUnavailable = SmtpCode{Code: 451, EnhancedCodes: [3]int16{}}

var InsufficientSystemStorage = SmtpCode{Code: 452, EnhancedCodes: NoEnhancedCodes}

var TemporaryAuthFail = SmtpCode{Code: 454, EnhancedCodes: [3]int16{}}

var ServerUnableToAccommodateParameters = SmtpCode{Code: 455, EnhancedCodes: NoEnhancedCodes}

var SyntaxError = SmtpCode{Code: 500, EnhancedCodes: NoEnhancedCodes}

var AuthExchangeLineIsTooLong = SmtpCode{Code: 500, EnhancedCodes: [3]int16{}}

var SyntaxErrorInParametersOrArguments = SmtpCode{Code: 501, EnhancedCodes: NoEnhancedCodes}

var CannotBase64DecodeClientResponses = SmtpCode{Code: 501, EnhancedCodes: [3]int16{}}

var ClientInitiatedAuthenticationExchange = SmtpCode{Code: 501, EnhancedCodes: [3]int16{}}

var CommandNotImplemented = SmtpCode{Code: 502, EnhancedCodes: NoEnhancedCodes}

var BadSequenceOfCommands = SmtpCode{Code: 503, EnhancedCodes: NoEnhancedCodes}

var CommandParameterIsNotImplemented = SmtpCode{Code: 504, EnhancedCodes: NoEnhancedCodes}

var UnrecognizedAuthenticationType = SmtpCode{Code: 504, EnhancedCodes: [3]int16{}}

var ServerDoesNotAcceptMail = SmtpCode{Code: 521, EnhancedCodes: NoEnhancedCodes}

var EncryptionNeeded = SmtpCode{Code: 523, EnhancedCodes: NoEnhancedCodes}

var AuthenticationRequired = SmtpCode{Code: 530, EnhancedCodes: [3]int16{}}

var AuthenticationMechanismIsWeak = SmtpCode{Code: 534, EnhancedCodes: [3]int16{}}

var AuthenticationCredentialsInvalid = SmtpCode{Code: 535, EnhancedCodes: [3]int16{}}

var EncryptionRequiredForRequestedAuthMechanism = SmtpCode{Code: 538, EnhancedCodes: [3]int16{}}

var MailboxUnavailable = SmtpCode{Code: 550, EnhancedCodes: NoEnhancedCodes}

var UserNotLocal = SmtpCode{Code: 551, EnhancedCodes: NoEnhancedCodes}

var ExceededStorageAllocation = SmtpCode{Code: 552, EnhancedCodes: NoEnhancedCodes}

var MailboxNameNotAllowed = SmtpCode{Code: 553, EnhancedCodes: NoEnhancedCodes}

var TransactionHasFailed = SmtpCode{Code: 554, EnhancedCodes: NoEnhancedCodes}

var NoSmtpServiceHere = SmtpCode{Code: 554, EnhancedCodes: NoEnhancedCodes}

var MessageTooBigForSystem = SmtpCode{Code: 554, EnhancedCodes: [3]int16{}}

var DomainDoesNotAcceptMail = SmtpCode{Code: 556, EnhancedCodes: NoEnhancedCodes}
