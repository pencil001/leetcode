var getIntersectionNode = function(headA, headB) {
    if(!headA || !headB){
        return null
    }
    
    let tailA = headA
    while(tailA.next) {
        tailA = tailA.next
    }
    tailA.next = headA
    
    let fast = headB
    let slow = headB
    let hasCycle = false
    
    while(fast && fast.next) {
        slow = slow.next
        fast = fast.next.next
        if(slow == fast) {
            hasCycle = true
            break
        }
    }
    
    let interNode = null
    if(hasCycle) {
        let p1 = headB
        let p2 = fast
        while (p1 != p2) {
            p1 = p1.next
            p2 = p2.next
        }
        interNode = p1
    }
    
    tailA.next = null
    return interNode
};

var getIntersectionNode2 = function(headA, headB) {
    let cur1 = headA;
    let cur2 = headB;
    let nextList1 = false;
    let nextList2 = false;

    while (cur1 !== null && cur2 !== null) {
        if(nextList1 === true && nextList2 === true) {
            if(cur1 === cur2){
                return cur1;
            } else if (cur1 === null || cur2 === null) {
                return null;
            }
        }
        
        cur1 = cur1.next;
        cur2 = cur2.next;

        if (!cur1 && !nextList1) {
            cur1 = headB;
            nextList1 = true;
        }
        if (!cur2 && !nextList2) {
            cur2 = headA;
            nextList2 = true;
        }

    }
    return null;
};