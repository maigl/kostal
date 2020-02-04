package data

// GENERATED FILE: see code_gen.bash
var Registers = []*Register{
    &Register{
        Addr: 2,
        Unit: "-",
        Format: "Bool",
        Length: 1,
        Description: "MODBUS Enable",
    },
    &Register{
        Addr: 4,
        Unit: "-",
        Format: "U16",
        Length: 1,
        Description: "MODBUS Unit-ID",
    },
    &Register{
        Addr: 6,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "Inverter article number",
    },
    &Register{
        Addr: 14,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "Inverter serial number",
    },
    &Register{
        Addr: 30,
        Unit: "-",
        Format: "U16",
        Length: 1,
        Description: "Number of bidirectional converter",
    },
    &Register{
        Addr: 32,
        Unit: "-",
        Format: "U16",
        Length: 1,
        Description: "Number of AC phases",
    },
    &Register{
        Addr: 34,
        Unit: "-",
        Format: "U16",
        Length: 1,
        Description: "Number of PV strings",
    },
    &Register{
        Addr: 36,
        Unit: "-",
        Format: "U16",
        Length: 2,
        Description: "Hardware-Version",
    },
    &Register{
        Addr: 38,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "Software-Version Maincontroller (MC)",
    },
    &Register{
        Addr: 46,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "Software-Version IO-Controller (IOC)",
    },
    &Register{
        Addr: 54,
        Unit: "-",
        Format: "U16",
        Length: 2,
        Description: "Power-ID",
    },
    &Register{
        Addr: 56,
        Unit: "-",
        Format: "U16",
        Length: 2,
        Description: "Inverter state2",
    },
    &Register{
        Addr: 100,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Total DC power",
    },
    &Register{
        Addr: 104,
        Unit: "-",
        Format: "U32",
        Length: 2,
        Description: "State of energy manager3",
    },
    &Register{
        Addr: 106,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Home own consumption from battery",
    },
    &Register{
        Addr: 108,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Home own consumption from grid",
    },
    &Register{
        Addr: 110,
        Unit: "Wh",
        Format: "Float",
        Length: 2,
        Description: "Total home consumption Battery",
    },
    &Register{
        Addr: 112,
        Unit: "Wh",
        Format: "Float",
        Length: 2,
        Description: "Total home consumption Grid",
    },
    &Register{
        Addr: 114,
        Unit: "Wh",
        Format: "Float",
        Length: 2,
        Description: "Total home consumption PV",
    },
    &Register{
        Addr: 116,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Home own consumption from PV",
    },
    &Register{
        Addr: 118,
        Unit: "Wh",
        Format: "Float",
        Length: 2,
        Description: "Total home consumption",
    },
    &Register{
        Addr: 120,
        Unit: "Ohm",
        Format: "Float",
        Length: 2,
        Description: "Isolation resistance",
    },
    &Register{
        Addr: 122,
        Unit: "%",
        Format: "Float",
        Length: 2,
        Description: "Power limit from EVU",
    },
    &Register{
        Addr: 124,
        Unit: "%",
        Format: "Float",
        Length: 2,
        Description: "Total home consumption rate",
    },
    &Register{
        Addr: 144,
        Unit: "s",
        Format: "Float",
        Length: 2,
        Description: "Worktime",
    },
    &Register{
        Addr: 150,
        Unit: "-",
        Format: "Float",
        Length: 2,
        Description: "Actual cos φ",
    },
    &Register{
        Addr: 152,
        Unit: "Hz",
        Format: "Float",
        Length: 2,
        Description: "Grid frequency",
    },
    &Register{
        Addr: 154,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Current Phase 1",
    },
    &Register{
        Addr: 156,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Active power Phase 1",
    },
    &Register{
        Addr: 158,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Voltage Phase 1",
    },
    &Register{
        Addr: 160,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Current Phase 2",
    },
    &Register{
        Addr: 162,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Active power Phase 2",
    },
    &Register{
        Addr: 164,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Voltage Phase 2",
    },
    &Register{
        Addr: 166,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Current Phase 3",
    },
    &Register{
        Addr: 168,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Active power Phase 3",
    },
    &Register{
        Addr: 170,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Voltage Phase 3",
    },
    &Register{
        Addr: 172,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Total AC active power",
    },
    &Register{
        Addr: 174,
        Unit: "Var",
        Format: "Float",
        Length: 2,
        Description: "Total AC reactive power",
    },
    &Register{
        Addr: 178,
        Unit: "VA",
        Format: "Float",
        Length: 2,
        Description: "Total AC apparent power",
    },
    &Register{
        Addr: 190,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Battery charge current",
    },
    &Register{
        Addr: 194,
        Unit: "-",
        Format: "Float",
        Length: 2,
        Description: "Number of battery cycles",
    },
    &Register{
        Addr: 200,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Actual battery charge (-) / discharge (+) current",
    },
    &Register{
        Addr: 202,
        Unit: "-",
        Format: "Float",
        Length: 2,
        Description: "PSSB fuse state5",
    },
    &Register{
        Addr: 208,
        Unit: "-",
        Format: "Float",
        Length: 2,
        Description: "Battery ready flag",
    },
    &Register{
        Addr: 210,
        Unit: "%",
        Format: "Float",
        Length: 2,
        Description: "Act. state of charge",
    },
    &Register{
        Addr: 214,
        Unit: "°C",
        Format: "Float",
        Length: 2,
        Description: "Battery temperature",
    },
    &Register{
        Addr: 216,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Battery voltage",
    },
    &Register{
        Addr: 218,
        Unit: "-",
        Format: "Float",
        Length: 2,
        Description: "Cos φ (powermeter)",
    },
    &Register{
        Addr: 220,
        Unit: "Hz",
        Format: "Float",
        Length: 2,
        Description: "Frequency (powermeter)",
    },
    &Register{
        Addr: 222,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Current phase 1 (powermeter)",
    },
    &Register{
        Addr: 224,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Active power phase 1 (powermeter)",
    },
    &Register{
        Addr: 226,
        Unit: "Var",
        Format: "Float",
        Length: 2,
        Description: "Reactive power phase 1 (powermeter)",
    },
    &Register{
        Addr: 228,
        Unit: "VA",
        Format: "Float",
        Length: 2,
        Description: "Apparent power phase 1 (powermeter)",
    },
    &Register{
        Addr: 230,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Voltage phase 1 (powermeter)",
    },
    &Register{
        Addr: 232,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Current phase 2 (powermeter)",
    },
    &Register{
        Addr: 234,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Active power phase 2 (powermeter)",
    },
    &Register{
        Addr: 236,
        Unit: "Var",
        Format: "Float",
        Length: 2,
        Description: "Reactive power phase 2 (powermeter)",
    },
    &Register{
        Addr: 238,
        Unit: "VA",
        Format: "Float",
        Length: 2,
        Description: "Apparent power phase 2 (powermeter)",
    },
    &Register{
        Addr: 240,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Voltage phase 2 (powermeter)",
    },
    &Register{
        Addr: 242,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Current phase 3 (powermeter)",
    },
    &Register{
        Addr: 244,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Active power phase 3 (powermeter)",
    },
    &Register{
        Addr: 246,
        Unit: "Var",
        Format: "Float",
        Length: 2,
        Description: "Reactive power phase 3 (powermeter)",
    },
    &Register{
        Addr: 248,
        Unit: "VA",
        Format: "Float",
        Length: 2,
        Description: "Apparent power phase 3 (powermeter)",
    },
    &Register{
        Addr: 250,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Voltage phase 3 (powermeter)",
    },
    &Register{
        Addr: 252,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Total active power (powermeter)",
    },
    &Register{
        Addr: 254,
        Unit: "Var",
        Format: "Float",
        Length: 2,
        Description: "Total reactive power (powermeter)",
    },
    &Register{
        Addr: 256,
        Unit: "VA",
        Format: "Float",
        Length: 2,
        Description: "Total apparent power (powermeter)",
    },
    &Register{
        Addr: 258,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Current DC1",
    },
    &Register{
        Addr: 260,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Power DC1",
    },
    &Register{
        Addr: 266,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Voltage DC1",
    },
    &Register{
        Addr: 268,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Current DC2",
    },
    &Register{
        Addr: 270,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Power DC2",
    },
    &Register{
        Addr: 276,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Voltage DC2",
    },
    &Register{
        Addr: 278,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Current DC3",
    },
    &Register{
        Addr: 280,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Power DC3",
    },
    &Register{
        Addr: 286,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Voltage DC3",
    },
    &Register{
        Addr: 320,
        Unit: "Wh",
        Format: "Float",
        Length: 2,
        Description: "Total yield",
    },
    &Register{
        Addr: 322,
        Unit: "Wh",
        Format: "Float",
        Length: 2,
        Description: "Daily yield",
    },
    &Register{
        Addr: 324,
        Unit: "Wh",
        Format: "Float",
        Length: 2,
        Description: "Yearly yield",
    },
    &Register{
        Addr: 326,
        Unit: "Wh",
        Format: "Float",
        Length: 2,
        Description: "Monthly yield",
    },
    &Register{
        Addr: 384,
        Unit: "-",
        Format: "String",
        Length: 32,
        Description: "Inverter network name",
    },
    &Register{
        Addr: 416,
        Unit: "-",
        Format: "U16",
        Length: 1,
        Description: "IP enable",
    },
    &Register{
        Addr: 418,
        Unit: "-",
        Format: "U16",
        Length: 1,
        Description: "Manual IP / Auto-IP",
    },
    &Register{
        Addr: 420,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "IP-address",
    },
    &Register{
        Addr: 428,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "IP-subnetmask",
    },
    &Register{
        Addr: 436,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "IP-gateway",
    },
    &Register{
        Addr: 444,
        Unit: "-",
        Format: "U16",
        Length: 1,
        Description: "IP-auto-DNS",
    },
    &Register{
        Addr: 446,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "IP-DNS1",
    },
    &Register{
        Addr: 454,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "IP-DNS2",
    },
    &Register{
        Addr: 512,
        Unit: "Ah",
        Format: "U32",
        Length: 2,
        Description: "Battery gross capacity",
    },
    &Register{
        Addr: 514,
        Unit: "%",
        Format: "U16",
        Length: 1,
        Description: "Battery actual SOC",
    },
    &Register{
        Addr: 515,
        Unit: "-",
        Format: "U32",
        Length: 2,
        Description: "Firmware Maincontroller (MC)",
    },
    &Register{
        Addr: 517,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "Battery Manufacturer",
    },
    &Register{
        Addr: 525,
        Unit: "-",
        Format: "U32",
        Length: 2,
        Description: "Battery Model ID",
    },
    &Register{
        Addr: 527,
        Unit: "-",
        Format: "U32",
        Length: 2,
        Description: "Battery Serial Number",
    },
    &Register{
        Addr: 529,
        Unit: "Wh",
        Format: "U32",
        Length: 2,
        Description: "Work Capacity",
    },
    &Register{
        Addr: 531,
        Unit: "W",
        Format: "U16",
        Length: 1,
        Description: "Inverter Max Power",
    },
    &Register{
        Addr: 532,
        Unit: "-",
        Format: "-",
        Length: 1,
        Description: "Inverter Peak Generation Power Scale Factor4",
    },
    &Register{
        Addr: 535,
        Unit: "-",
        Format: "String",
        Length: 16,
        Description: "Inverter Manufacturer",
    },
    &Register{
        Addr: 551,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "Inverter Model ID",
    },
    &Register{
        Addr: 559,
        Unit: "-",
        Format: "String",
        Length: 16,
        Description: "Inverter Serial Number",
    },
    &Register{
        Addr: 575,
        Unit: "W",
        Format: "S16",
        Length: 1,
        Description: "Inverter Generation Power (actual)",
    },
    &Register{
        Addr: 576,
        Unit: "-",
        Format: "-",
        Length: 1,
        Description: "Power Scale Factor4",
    },
    &Register{
        Addr: 577,
        Unit: "Wh",
        Format: "U32",
        Length: 2,
        Description: "Generation Energy",
    },
    &Register{
        Addr: 579,
        Unit: "-",
        Format: "-",
        Length: 1,
        Description: "Energy Scale Factor4",
    },
    &Register{
        Addr: 582,
        Unit: "W",
        Format: "S16",
        Length: 1,
        Description: "Actual battery charge/discharge power",
    },
    &Register{
        Addr: 586,
        Unit: "-",
        Format: "U16",
        Length: 1,
        Description: "Battery Firmware",
    },
    &Register{
        Addr: 588,
        Unit: "-",
        Format: "U16",
        Length: 1,
        Description: "Battery Type6",
    },
    &Register{
        Addr: 768,
        Unit: "-",
        Format: "String",
        Length: 32,
        Description: "Productname (e.g. PLENTICORE plus)",
    },
    &Register{
        Addr: 800,
        Unit: "-",
        Format: "String",
        Length: 32,
        Description: "Power class (e.g. 10)",
    },
}