public class Main {
	
	public static void main(String[] args) {
		
		// Uniformly distributed values
		int[] numbers = {2, 5, 8, 11, 14, 17, 20, 23};
		int find = 17;
		System.out.println("Position: " + InterpolationSearch.search(numbers, find));
		
		// Non-Uniformly distributed values
		int[] moreNumbers = {1, 3, 9, 10, 18, 20, 25, 40, 50};
		int findMore = 25;
		System.out.println("Position: " + InterpolationSearch.search(moreNumbers, findMore));
	}
}
