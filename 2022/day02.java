import java.io.File;
import java.io.FileNotFoundException;
import java.util.HashMap;
import java.util.Scanner;

public class day02 {
    public static void main(String[] args) {
        long startTime = System.nanoTime();
        HashMap<String, Integer> points = new HashMap<String, Integer>();
        points.put("X", 1);
        points.put("Y", 2);
        points.put("Z", 3);
        points.put("A", 1);
        points.put("B", 2);
        points.put("C", 3);
        int firstPartCounter = 0, secondPartCounter = 0;
        try {
            Scanner myReader = new Scanner(new File("input.txt"));
            while (myReader.hasNextLine()) {
                String[] plays = myReader.nextLine().split(" ");
                String op = plays[0];
                String mine = plays[1];

                if((int)op.charAt(0) == (int)mine.charAt(0)-23){
                    firstPartCounter += 3;
                }else if(   (op.equalsIgnoreCase("C") && mine.equalsIgnoreCase("X")) || 
                            (op.equalsIgnoreCase("A") && mine.equalsIgnoreCase("Y")) || 
                            (op.equalsIgnoreCase("B") && mine.equalsIgnoreCase("Z")))   {
                    firstPartCounter += 6;
                }

                switch(mine){
                    case "X": //Lose
                        if(op.equalsIgnoreCase("A")) secondPartCounter += 3;
                        else if(op.equalsIgnoreCase("B")) secondPartCounter += 1;
                        else secondPartCounter += 2;
                    break;

                    case "Y": //Draw
                        secondPartCounter += points.get(op) + 3;
                    break;

                    case "Z": //Win
                        if(op.equalsIgnoreCase("A")) secondPartCounter += 2;
                        else if(op.equalsIgnoreCase("B")) secondPartCounter += 3;
                        else secondPartCounter += 1;
                        secondPartCounter += 6;
                    break;

                    default: throw new IllegalArgumentException("Input corrupted!");
                }


                firstPartCounter += points.get(mine);
            }
            myReader.close();
        } catch (FileNotFoundException e) {
            System.out.println("Invalid input file.");
            e.printStackTrace();
        }
        System.out.println("FIRST PART: "+firstPartCounter);
        System.out.println("SECOND PART: "+secondPartCounter);
        System.out.println("Execution time: "+(System.nanoTime()-startTime)/1000000 + " millis");
    }

}
