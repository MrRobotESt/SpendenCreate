# SpendenCreate

## Control Sequenz 


1. If you using the Programm for the first Time, start csvconvert!


    var csvConverter = flag.Bool("csvconvert", false, "csv has to be convert in german iso8859-1 Chars")

2. For the whole PDF creater process, you have to use start!


    var startCreator = flag.Bool("start", false, "start the pdf creator")

3. To check if some person has made a donation but it's not in the Adress CSV

    
    var checkexistence = flag.Bool("existence", false, "check the existence of Person in DepositData")
