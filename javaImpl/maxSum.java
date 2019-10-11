//John miner
//package maxSum;
import java.math.*;
import java.util.*;

class maxSum {

	static int stupidFreshmen(int n, ArrayList<Integer> a) {

		int maxSum = 0;
		// i: starting index of sum
		for (int i = 0; i < n; i++) {
			// j: ending index of sum
			for (int j = i; j < n; j++) {
				int thisSum = 0;
				for (int k = i; k <= j; k++) {
					thisSum += a.get(k);
				}
				if (thisSum > maxSum)
					maxSum = thisSum;
			}
		}
		return maxSum;
	}

	static int slightlyLessStupidSophomore(int n, ArrayList<Integer> a) {
		int maxSum = 0;
		// i: starting index of sum
		for (int i = 0; i < n; i++) {
			int thisSum = 0;
			// compute all sums that begin with index i
			for (int j = i; j < n; j++) {
				thisSum += a.get(j);
				if (thisSum > maxSum)
					maxSum = thisSum;
			}
		}
		return maxSum;
	}

	static int gettingSmarterJunior(ArrayList<Integer> a, int left, int right) {
		// Base case 1
		if (right == left)
			return a.get(left);
		// Base case 2
		if (right == left + 1)
			return Math.max(a.get(left), Math.max(a.get(right), a.get(left) + a.get(right)));
		int mid = (left + right) / 2;
		// Find the MSS that occurs in the left half of a
		int mss_left = gettingSmarterJunior(a, left, mid);
		// Find the MSS that occurs in the right half of a
		int mss_right = gettingSmarterJunior(a, mid + 1, right);
		// Find the MSS that intersects both the left and right halves
		int mss_middle = gettingSmarterJuniorMiddle(a, left, mid, right);
		return Math.max(mss_left, Math.max(mss_right, mss_middle));
	}

	static int gettingSmarterJuniorMiddle(ArrayList<Integer> a, int left, int mid, int right) {
		int negativeInfinity = Integer.MIN_VALUE;
		int max_left_sum = negativeInfinity;
		int sum = 0;
		int i;
		for (i = mid; i >= left; i--) {
			sum += a.get(i);
			if (sum > max_left_sum)
				max_left_sum = sum;
		}
		int max_right_sum = negativeInfinity;
		sum = 0;
		for (i = mid + 1; i <= right; i++) {
			sum += a.get(i);
			if (sum > max_right_sum)
				max_right_sum = sum;
		}
		return max_left_sum + max_right_sum;
	}

	static int superSmartSenior(ArrayList<Integer> a) {
		int maxSum = 0;
		int thisSum = 0;
		for (int i = 0; i < a.size(); i++) {
			thisSum += a.get(i);
			if (thisSum > maxSum)
				maxSum = thisSum;
			else if (thisSum < 0)
				thisSum = 0;
		}
		return maxSum;
	}

	public static void main(String[] args) {

		ArrayList<Integer> array = new ArrayList<Integer>();
		ArrayList<Integer> arraySecond = new ArrayList<Integer>();
		Scanner in = new Scanner(System.in);

		int selection = 0;

		while (selection != 4) {
			System.out.println("Please make a selection:\n1) Make your own array \n"
					+ "2) Time a random array \n3) Select one of the algorithms for computing \n4) Quit ");
			selection = in.nextInt();
			switch (selection) {
			case 1: {
				System.out.println("Enter each element of your array followed by commas:");
				String reply = in.next();
				String[] values = reply.split(",");
				for (String var : values) {
					array.add(Integer.parseInt(var));
				}
				int n = array.size();
				int l = 0;
				int r = array.size() - 1;
				int algorithmOne = stupidFreshmen(n, array);
				int algorithmTwo = slightlyLessStupidSophomore(n, array);
				int algorithmThree = gettingSmarterJunior(array, l, r);
				int algorithmFour = superSmartSenior(array);

				System.out.println("Your stupid sum 1 is " + algorithmOne);
				System.out.println("Your stupid sum 2 is " + algorithmTwo);
				System.out.println("Your stupid sum 3 is " + algorithmThree);
				System.out.println("Your stupid sum 4 is " + algorithmFour);
			}
				break;

			case 2: {
				System.out.println("Enter number of elements you want in the array");
				int n = in.nextInt();
				int i = 0;
				Random rand = new Random();
				System.out.println("Enter a value for your array ");

				for (i = 0; i < n; i++) {
					int value = rand.nextInt(((100 - 0) + 1) - 50);
					array.add(value);
				}
				int l = 0;
				int r = array.size() - 1;

				{

					System.out.println("Which algorithms would you like to run?\n Choose 1-4");
					String algorithmChoice = in.next();
					char choiceArray[] = algorithmChoice.toCharArray();

					int c = 0;

					while (c != choiceArray.length) {
						System.out.println(choiceArray[c]);
						switch (choiceArray[c]) {

						case '1': {

							long start = System.currentTimeMillis();
							int algorithmOne = stupidFreshmen(n, array);
							long end = System.currentTimeMillis();
							System.out.println("Your time is " + (end - start) + " milli seconds");
							System.out.println("Your stupid sum 1 is " + algorithmOne);
						}
							break;

						case '2': {
							long start = System.currentTimeMillis();
							int algorithmTwo = slightlyLessStupidSophomore(n, array);
							long end = System.currentTimeMillis();
							System.out.println("Your time is " + (end - start) + " milli seconds");
							System.out.println("Your stupid sum 2 is " + algorithmTwo);
						}
							break;

						case '3': {
							long start = System.nanoTime();
							int algorithmThree = gettingSmarterJunior(array, l, r);
							long end = System.nanoTime();
							System.out.println("Your time is " + ((end - start)) + " nano seconds!!!");
							System.out.println("Your stupid sum 3 is " + algorithmThree);
						}
							break;

						case '4': {
							long start = System.nanoTime();
							int algorithmFour = superSmartSenior(array);
							long end = System.nanoTime();
							System.out.println("Your time is " + (end - start) + " nano seconds!!!");
							System.out.println("Your stupid sum 4 is " + algorithmFour);
						}
							break;
						}
						c++;
					}
				}

			}
				break;

			case 3: {

				System.out.println("Enter number of elements you want in the first array");
				int m = in.nextInt();
				int i = 0;

				Random rand = new Random();

				for (i = 0; i < m; i++) {
					int value = rand.nextInt(((100 - 0) + 1) - 50);
					array.add(value);

				}
				int l = 0;
				int r = array.size() - 1;
				/**
				 * second random generator
				 */
				System.out.println("How many elements in your other array?\n ");
				int n = in.nextInt();

				for (i = 0; i < n; i++) {
					int value = rand.nextInt(((100 - 0) + 1) - 50);
					arraySecond.add(value);

				}
				int lSecond = 0;
				int rSecond = arraySecond.size() - 1;
				{

					System.out.println("Which algorithm would you like to run?\n Choose 1-4");
					int choice = in.nextInt();

					switch (choice) {

					case 1: {
						long start = System.currentTimeMillis();
						int algorithmOne = stupidFreshmen(m, array);
						long end1 = System.currentTimeMillis();
						System.out.println((end1 - start));
						long prediction = (long) ((end1 - start) * (Math.pow(n, 3) / Math.pow(m, 3)));
						System.out.println("Your time should be: " + prediction);

						start = System.currentTimeMillis();
						algorithmOne = stupidFreshmen(n, arraySecond);
						long end = System.currentTimeMillis();
						System.out.println("Your time is " + (end - start) + " milli seconds");

					}
						break;

					case 2: {
						long start = System.currentTimeMillis();
						int algorithmTwo = slightlyLessStupidSophomore(m, array);
						long end1 = System.currentTimeMillis();

						long prediction = ((end1 - start) * (long) (Math.pow(n, 2) / Math.pow(m, 2)));
						System.out.println("Your time should be: " + prediction);

						start = System.currentTimeMillis();
						algorithmTwo = slightlyLessStupidSophomore(n, arraySecond);
						long end = System.currentTimeMillis();
						System.out.println("Your actual time is " + (end - start) + " milli seconds");

					}
						break;

					case 3: {
						long start = System.nanoTime();
						int algorithmThree = gettingSmarterJunior(array, l, r);
						long end1 = System.nanoTime();

						long prediction = ((end1 - start)
								* (long) ((Math.log(n) / Math.log(2)) / (Math.log(m) / Math.log(2))));
						System.out.println("Your time should be: " + prediction);

						start = System.nanoTime();
						algorithmThree = gettingSmarterJunior(arraySecond, lSecond, rSecond);
						long end = System.nanoTime();
						System.out.println("Your actual time is " + ((end - start)) + " nano seconds!!!");

					}
						break;

					case 4: {
						long start = System.nanoTime();
						int algorithmFour = superSmartSenior(array);
						long end1 = System.nanoTime();

						long prediction = ((end1 - start) * (long) (n / m));
						System.out.println("Your time should be: " + prediction);

						start = System.nanoTime();
						algorithmFour = superSmartSenior(arraySecond);
						long end = System.nanoTime();
						System.out.println("Your time is " + (end - start) + " nano seconds!!!");

					}
						break;
					}

				}

			}
				break;

			case 4: {
				System.out.println("Bye!");
			}
				break;
			}
		}
	}

}