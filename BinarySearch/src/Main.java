import java.util.*;

public class Main {
	
	public static void main(String[] args) {
		int[] numbers = {1, 2, 3, 4, 5};
		int find = 4;
		
		System.out.println(BinarySearch.recursive(0, 4, numbers, find));
		System.out.println(BinarySearch.iterative(numbers, find));
	}
}
