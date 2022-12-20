use std::borrow::Borrow;
use std::borrow::BorrowMut;
use std::option;
use std::mem;

use std::ptr;
use std::rc::Rc;
// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val:val,
    }
  }
}

fn main() {
    
    let list0:Option<Box<ListNode>>=None;
    let list1=Some(Box::new(ListNode{
        val:1,
        next:Some(Box::new(ListNode{
            val:2,
            next:None,
        })),
    }));
    let list2=Some(Box::new(ListNode{
        val:1,
        next:Some(Box::new(ListNode{
            val:2,
            next:Some(Box::new(ListNode{
                val:2,
                next:Some(Box::new(ListNode{
                    val:3,
                    next:None,
                })),
            })),
        })),
    }));

    // delete_duplicates(list0);
    // delete_duplicates(list2);
    



}

pub fn delete_duplicates(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    if head.is_none(){
      return None;
    }
    let mut node=head.as_mut().unwrap();
    
    while let Some(mut fast)=std::mem::replace(&mut node.next, None){
      
      if fast.val==node.val {
          node.next=fast.next;
      }else {
          node.next=Some(fast);
          node=node.next.as_mut().unwrap();
      }
    }


    return None;
}
