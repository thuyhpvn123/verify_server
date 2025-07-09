window.env = {
  contractAddress: "0xCc3049c411106398d22bd23869796500a0d14B84",
  contractABI: [
    {
      inputs: [
        {
          internalType: "string",
          name: "_phoneNumber",
          type: "string",
        },
        {
          internalType: "enum AuthOTP.TypeMethod",
          name: "_typeMethod",
          type: "uint8",
        },
      ],
      name: "addBot",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    {
      inputs: [
        {
          internalType: "string",
          name: "_data",
          type: "string",
        },
        {
          internalType: "string",
          name: "_publicKey",
          type: "string",
        },
      ],
      name: "completeAuthentication",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    {
      inputs: [],
      stateMutability: "nonpayable",
      type: "constructor",
    },
    {
      anonymous: false,
      inputs: [
        {
          indexed: true,
          internalType: "address",
          name: "user",
          type: "address",
        },
        {
          indexed: false,
          internalType: "uint256",
          name: "otp",
          type: "uint256",
        },
        {
          indexed: false,
          internalType: "string",
          name: "chatbotPhone",
          type: "string",
        },
        {
          indexed: false,
          internalType: "enum AuthOTP.TypeMethod",
          name: "typeMethod",
          type: "uint8",
        },
      ],
      name: "AuthenticationRequested",
      type: "event",
    },
    {
      inputs: [
        {
          internalType: "string",
          name: "_userPhoneNumber",
          type: "string",
        },
        {
          internalType: "string",
          name: "_publicKey",
          type: "string",
        },
        {
          internalType: "enum AuthOTP.TypeMethod",
          name: "_typeMethod",
          type: "uint8",
        },
      ],
      name: "requestAuthentication",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    {
      inputs: [
        {
          internalType: "uint256",
          name: "_botId",
          type: "uint256",
        },
        {
          internalType: "string",
          name: "_phoneNumber",
          type: "string",
        },
        {
          internalType: "enum AuthOTP.TypeMethod",
          name: "_typeMethod",
          type: "uint8",
        },
        {
          internalType: "bool",
          name: "_status",
          type: "bool",
        },
      ],
      name: "updateBot",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    {
      inputs: [
        {
          internalType: "uint256",
          name: "_otp",
          type: "uint256",
        },
        {
          internalType: "string",
          name: "userPhoneNumber",
          type: "string",
        },
      ],
      name: "validateOTP",
      outputs: [
        {
          internalType: "string",
          name: "",
          type: "string",
        },
      ],
      stateMutability: "nonpayable",
      type: "function",
    },
    {
      inputs: [
        {
          internalType: "uint256",
          name: "",
          type: "uint256",
        },
      ],
      name: "detailBots",
      outputs: [
        {
          internalType: "string",
          name: "phoneNumber",
          type: "string",
        },
        {
          internalType: "enum AuthOTP.TypeMethod",
          name: "typeMethod",
          type: "uint8",
        },
        {
          internalType: "bool",
          name: "busy",
          type: "bool",
        },
        {
          internalType: "uint256",
          name: "timeOccupied",
          type: "uint256",
        },
        {
          internalType: "bool",
          name: "status",
          type: "bool",
        },
      ],
      stateMutability: "view",
      type: "function",
    },
    {
      inputs: [
        {
          internalType: "string",
          name: "",
          type: "string",
        },
      ],
      name: "publicKeyUsers",
      outputs: [
        {
          internalType: "address",
          name: "",
          type: "address",
        },
      ],
      stateMutability: "view",
      type: "function",
    },
    {
      inputs: [
        {
          internalType: "string",
          name: "_publicKey",
          type: "string",
        },
        {
          internalType: "bytes32",
          name: "_dataHash",
          type: "bytes32",
        },
      ],
      name: "verifyHash",
      outputs: [
        {
          internalType: "bool",
          name: "",
          type: "bool",
        },
      ],
      stateMutability: "view",
      type: "function",
    },
  ],
};