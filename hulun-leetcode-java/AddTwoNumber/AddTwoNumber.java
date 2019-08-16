import org.w3c.dom.NodeList;

public class AddTwoNumber{
    /*
2 -> 4 -> 3
5 -> 6 -> 4
7    0    8
*/
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode fake = new ListNode(0);
        ListNode p = fake;

        ListNode p1 = l1;
        ListNode p2 = l2;

        int carry = 0;
        while(p1!=null || p2!=null || carry > 0){
            int sum = carry;
            if(p1!=null){
                sum += p1.val;
                p1 = p1.next;
            }

            if(p2!=null){
                sum += p2.val;
                p2 = p2.next;
            }

            carry = sum / 10;

            ListNode l = new ListNode(sum % 10);
            p.next = l;
            p = p.next;
        }
        return fake.next;
    }

    public boolean CompareLists(ListNode l1,ListNode l2){
        while (l1 != null && l2 != null && l1.val == l2.val){
            l1 = l1.next;
            l2 = l2.next;
        }

        if (l1   == null && l2 == null){
            return true;
        }
        return false;
    }

    public ListNode MakeListNode(int[] arr){
        ListNode listNode = new ListNode(arr[0]);
        ListNode p = listNode;
        for (int i = 1;i < arr.length;i++){
            p.next = new ListNode(arr[i]);
            p = p.next;
        }
        return listNode;
    }

    public static void main(String[] args){
        AddTwoNumber addTwoNumber = new AddTwoNumber();
        int arr1[] = new int[]{1,3,4,5};
        int arr2[] = new int[]{4,5,6};
        int arr3[] = new int[]{5,8,0,6};
        ListNode ln1 = addTwoNumber.MakeListNode(arr1);
        ListNode ln2 = addTwoNumber.MakeListNode(arr2);
        ListNode ln3 = addTwoNumber.MakeListNode(arr3);

        ListNode ans = addTwoNumber.addTwoNumbers(ln1,ln2);

        Boolean bool = addTwoNumber.CompareLists(ans,ln3);
        System.out.println(bool);

    }
}

