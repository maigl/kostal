package data

// GENERATED FILE: see code_gen.bash
var Registers = map[string]*Register{
    "2": &Register{
        Addr: 2,
        Unit: "-",
        Format: "Bool",
        Length: 1,
        Description: "MODBUS Enable",
    },
    "4": &Register{
        Addr: 4,
        Unit: "-",
        Format: "U16",
        Length: 1,
        Description: "MODBUS Unit-ID",
    },
    "6": &Register{
        Addr: 6,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "Inverter article number",
    },
    "14": &Register{
        Addr: 14,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "Inverter serial number",
    },
    "30": &Register{
        Addr: 30,
        Unit: "-",
        Format: "U16",
        Length: 1,
        Description: "Number of bidirectional converter",
    },
    "32": &Register{
        Addr: 32,
        Unit: "-",
        Format: "U16",
        Length: 1,
        Description: "Number of AC phases",
    },
    "34": &Register{
        Addr: 34,
        Unit: "-",
        Format: "U16",
        Length: 1,
        Description: "Number of PV strings",
    },
    "36": &Register{
        Addr: 36,
        Unit: "-",
        Format: "U16",
        Length: 2,
        Description: "Hardware-Version",
    },
    "38": &Register{
        Addr: 38,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "Software-Version Maincontroller (MC)",
    },
    "46": &Register{
        Addr: 46,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "Software-Version IO-Controller (IOC)",
    },
    "54": &Register{
        Addr: 54,
        Unit: "-",
        Format: "U16",
        Length: 2,
        Description: "Power-ID",
    },
    "56": &Register{
        Addr: 56,
        Unit: "-",
        Format: "U16",
        Length: 2,
        Description: "Inverter state2",
    },
    "100": &Register{
        Addr: 100,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Total DC power",
    },
    "104": &Register{
        Addr: 104,
        Unit: "-",
        Format: "U32",
        Length: 2,
        Description: "State of energy manager3",
    },
    "106": &Register{
        Addr: 106,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Home own consumption from battery",
    },
    "108": &Register{
        Addr: 108,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Home own consumption from grid",
    },
    "110": &Register{
        Addr: 110,
        Unit: "Wh",
        Format: "Float",
        Length: 2,
        Description: "Total home consumption Battery",
    },
    "112": &Register{
        Addr: 112,
        Unit: "Wh",
        Format: "Float",
        Length: 2,
        Description: "Total home consumption Grid",
    },
    "114": &Register{
        Addr: 114,
        Unit: "Wh",
        Format: "Float",
        Length: 2,
        Description: "Total home consumption PV",
    },
    "116": &Register{
        Addr: 116,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Home own consumption from PV",
    },
    "118": &Register{
        Addr: 118,
        Unit: "Wh",
        Format: "Float",
        Length: 2,
        Description: "Total home consumption",
    },
    "120": &Register{
        Addr: 120,
        Unit: "Ohm",
        Format: "Float",
        Length: 2,
        Description: "Isolation resistance",
    },
    "122": &Register{
        Addr: 122,
        Unit: "%",
        Format: "Float",
        Length: 2,
        Description: "Power limit from EVU",
    },
    "124": &Register{
        Addr: 124,
        Unit: "%",
        Format: "Float",
        Length: 2,
        Description: "Total home consumption rate",
    },
    "144": &Register{
        Addr: 144,
        Unit: "s",
        Format: "Float",
        Length: 2,
        Description: "Worktime",
    },
    "150": &Register{
        Addr: 150,
        Unit: "-",
        Format: "Float",
        Length: 2,
        Description: "Actual cos φ",
    },
    "152": &Register{
        Addr: 152,
        Unit: "Hz",
        Format: "Float",
        Length: 2,
        Description: "Grid frequency",
    },
    "154": &Register{
        Addr: 154,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Current Phase 1",
    },
    "156": &Register{
        Addr: 156,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Active power Phase 1",
    },
    "158": &Register{
        Addr: 158,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Voltage Phase 1",
    },
    "160": &Register{
        Addr: 160,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Current Phase 2",
    },
    "162": &Register{
        Addr: 162,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Active power Phase 2",
    },
    "164": &Register{
        Addr: 164,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Voltage Phase 2",
    },
    "166": &Register{
        Addr: 166,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Current Phase 3",
    },
    "168": &Register{
        Addr: 168,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Active power Phase 3",
    },
    "170": &Register{
        Addr: 170,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Voltage Phase 3",
    },
    "172": &Register{
        Addr: 172,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Total AC active power",
    },
    "174": &Register{
        Addr: 174,
        Unit: "Var",
        Format: "Float",
        Length: 2,
        Description: "Total AC reactive power",
    },
    "178": &Register{
        Addr: 178,
        Unit: "VA",
        Format: "Float",
        Length: 2,
        Description: "Total AC apparent power",
    },
    "190": &Register{
        Addr: 190,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Battery charge current",
    },
    "194": &Register{
        Addr: 194,
        Unit: "-",
        Format: "Float",
        Length: 2,
        Description: "Number of battery cycles",
    },
    "200": &Register{
        Addr: 200,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Actual battery charge (-) / discharge (+) current",
    },
    "202": &Register{
        Addr: 202,
        Unit: "-",
        Format: "Float",
        Length: 2,
        Description: "PSSB fuse state5",
    },
    "208": &Register{
        Addr: 208,
        Unit: "-",
        Format: "Float",
        Length: 2,
        Description: "Battery ready flag",
    },
    "210": &Register{
        Addr: 210,
        Unit: "%",
        Format: "Float",
        Length: 2,
        Description: "Act. state of charge",
    },
    "214": &Register{
        Addr: 214,
        Unit: "°C",
        Format: "Float",
        Length: 2,
        Description: "Battery temperature",
    },
    "216": &Register{
        Addr: 216,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Battery voltage",
    },
    "218": &Register{
        Addr: 218,
        Unit: "-",
        Format: "Float",
        Length: 2,
        Description: "Cos φ (powermeter)",
    },
    "220": &Register{
        Addr: 220,
        Unit: "Hz",
        Format: "Float",
        Length: 2,
        Description: "Frequency (powermeter)",
    },
    "222": &Register{
        Addr: 222,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Current phase 1 (powermeter)",
    },
    "224": &Register{
        Addr: 224,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Active power phase 1 (powermeter)",
    },
    "226": &Register{
        Addr: 226,
        Unit: "Var",
        Format: "Float",
        Length: 2,
        Description: "Reactive power phase 1 (powermeter)",
    },
    "228": &Register{
        Addr: 228,
        Unit: "VA",
        Format: "Float",
        Length: 2,
        Description: "Apparent power phase 1 (powermeter)",
    },
    "230": &Register{
        Addr: 230,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Voltage phase 1 (powermeter)",
    },
    "232": &Register{
        Addr: 232,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Current phase 2 (powermeter)",
    },
    "234": &Register{
        Addr: 234,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Active power phase 2 (powermeter)",
    },
    "236": &Register{
        Addr: 236,
        Unit: "Var",
        Format: "Float",
        Length: 2,
        Description: "Reactive power phase 2 (powermeter)",
    },
    "238": &Register{
        Addr: 238,
        Unit: "VA",
        Format: "Float",
        Length: 2,
        Description: "Apparent power phase 2 (powermeter)",
    },
    "240": &Register{
        Addr: 240,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Voltage phase 2 (powermeter)",
    },
    "242": &Register{
        Addr: 242,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Current phase 3 (powermeter)",
    },
    "244": &Register{
        Addr: 244,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Active power phase 3 (powermeter)",
    },
    "246": &Register{
        Addr: 246,
        Unit: "Var",
        Format: "Float",
        Length: 2,
        Description: "Reactive power phase 3 (powermeter)",
    },
    "248": &Register{
        Addr: 248,
        Unit: "VA",
        Format: "Float",
        Length: 2,
        Description: "Apparent power phase 3 (powermeter)",
    },
    "250": &Register{
        Addr: 250,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Voltage phase 3 (powermeter)",
    },
    "252": &Register{
        Addr: 252,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Total active power (powermeter)",
    },
    "254": &Register{
        Addr: 254,
        Unit: "Var",
        Format: "Float",
        Length: 2,
        Description: "Total reactive power (powermeter)",
    },
    "256": &Register{
        Addr: 256,
        Unit: "VA",
        Format: "Float",
        Length: 2,
        Description: "Total apparent power (powermeter)",
    },
    "258": &Register{
        Addr: 258,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Current DC1",
    },
    "260": &Register{
        Addr: 260,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Power DC1",
    },
    "266": &Register{
        Addr: 266,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Voltage DC1",
    },
    "268": &Register{
        Addr: 268,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Current DC2",
    },
    "270": &Register{
        Addr: 270,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Power DC2",
    },
    "276": &Register{
        Addr: 276,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Voltage DC2",
    },
    "278": &Register{
        Addr: 278,
        Unit: "A",
        Format: "Float",
        Length: 2,
        Description: "Current DC3",
    },
    "280": &Register{
        Addr: 280,
        Unit: "W",
        Format: "Float",
        Length: 2,
        Description: "Power DC3",
    },
    "286": &Register{
        Addr: 286,
        Unit: "V",
        Format: "Float",
        Length: 2,
        Description: "Voltage DC3",
    },
    "320": &Register{
        Addr: 320,
        Unit: "Wh",
        Format: "Float",
        Length: 2,
        Description: "Total yield",
    },
    "322": &Register{
        Addr: 322,
        Unit: "Wh",
        Format: "Float",
        Length: 2,
        Description: "Daily yield",
    },
    "324": &Register{
        Addr: 324,
        Unit: "Wh",
        Format: "Float",
        Length: 2,
        Description: "Yearly yield",
    },
    "326": &Register{
        Addr: 326,
        Unit: "Wh",
        Format: "Float",
        Length: 2,
        Description: "Monthly yield",
    },
    "384": &Register{
        Addr: 384,
        Unit: "-",
        Format: "String",
        Length: 32,
        Description: "Inverter network name",
    },
    "416": &Register{
        Addr: 416,
        Unit: "-",
        Format: "U16",
        Length: 1,
        Description: "IP enable",
    },
    "418": &Register{
        Addr: 418,
        Unit: "-",
        Format: "U16",
        Length: 1,
        Description: "Manual IP / Auto-IP",
    },
    "420": &Register{
        Addr: 420,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "IP-address",
    },
    "428": &Register{
        Addr: 428,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "IP-subnetmask",
    },
    "436": &Register{
        Addr: 436,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "IP-gateway",
    },
    "444": &Register{
        Addr: 444,
        Unit: "-",
        Format: "U16",
        Length: 1,
        Description: "IP-auto-DNS",
    },
    "446": &Register{
        Addr: 446,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "IP-DNS1",
    },
    "454": &Register{
        Addr: 454,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "IP-DNS2",
    },
    "512": &Register{
        Addr: 512,
        Unit: "Ah",
        Format: "U32",
        Length: 2,
        Description: "Battery gross capacity",
    },
    "514": &Register{
        Addr: 514,
        Unit: "%",
        Format: "U16",
        Length: 1,
        Description: "Battery actual SOC",
    },
    "515": &Register{
        Addr: 515,
        Unit: "-",
        Format: "U32",
        Length: 2,
        Description: "Firmware Maincontroller (MC)",
    },
    "517": &Register{
        Addr: 517,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "Battery Manufacturer",
    },
    "525": &Register{
        Addr: 525,
        Unit: "-",
        Format: "U32",
        Length: 2,
        Description: "Battery Model ID",
    },
    "527": &Register{
        Addr: 527,
        Unit: "-",
        Format: "U32",
        Length: 2,
        Description: "Battery Serial Number",
    },
    "529": &Register{
        Addr: 529,
        Unit: "Wh",
        Format: "U32",
        Length: 2,
        Description: "Work Capacity",
    },
    "531": &Register{
        Addr: 531,
        Unit: "W",
        Format: "U16",
        Length: 1,
        Description: "Inverter Max Power",
    },
    "532": &Register{
        Addr: 532,
        Unit: "-",
        Format: "-",
        Length: 1,
        Description: "Inverter Peak Generation Power Scale Factor4",
    },
    "535": &Register{
        Addr: 535,
        Unit: "-",
        Format: "String",
        Length: 16,
        Description: "Inverter Manufacturer",
    },
    "551": &Register{
        Addr: 551,
        Unit: "-",
        Format: "String",
        Length: 8,
        Description: "Inverter Model ID",
    },
    "559": &Register{
        Addr: 559,
        Unit: "-",
        Format: "String",
        Length: 16,
        Description: "Inverter Serial Number",
    },
    "575": &Register{
        Addr: 575,
        Unit: "W",
        Format: "S16",
        Length: 1,
        Description: "Inverter Generation Power (actual)",
    },
    "576": &Register{
        Addr: 576,
        Unit: "-",
        Format: "-",
        Length: 1,
        Description: "Power Scale Factor4",
    },
    "577": &Register{
        Addr: 577,
        Unit: "Wh",
        Format: "U32",
        Length: 2,
        Description: "Generation Energy",
    },
    "579": &Register{
        Addr: 579,
        Unit: "-",
        Format: "-",
        Length: 1,
        Description: "Energy Scale Factor4",
    },
    "582": &Register{
        Addr: 582,
        Unit: "W",
        Format: "S16",
        Length: 1,
        Description: "Actual battery charge/discharge power",
    },
    "586": &Register{
        Addr: 586,
        Unit: "-",
        Format: "U16",
        Length: 1,
        Description: "Battery Firmware",
    },
    "588": &Register{
        Addr: 588,
        Unit: "-",
        Format: "U16",
        Length: 1,
        Description: "Battery Type6",
    },
    "768": &Register{
        Addr: 768,
        Unit: "-",
        Format: "String",
        Length: 32,
        Description: "Productname (e.g. PLENTICORE plus)",
    },
    "800": &Register{
        Addr: 800,
        Unit: "-",
        Format: "String",
        Length: 32,
        Description: "Power class (e.g. 10)",
    },
}
