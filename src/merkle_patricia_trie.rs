extern crate crypto;
extern crate rustc_serialize;

use std::collections::HashMap;
use self::crypto::digest::Digest;
use self::crypto::sha3::Sha3;
use rustc_serialize::hex::ToHex;
use std::u8;

enum Node {
    Branch([String; 17]),
    Flag((Vec<u8>, String)),
    Null(),
}

struct MerklePatriciaTrie {
    db: HashMap<String, Node>,
    root: String,
}

impl MerklePatriciaTrie {
    // fn get(&mut self, key: &str) -> String {
    //     // TODO
    // }

    // fn insert(&mut self, key: &str, new_value: &str) {
    //     // TODO
    //     if self.db.is_empty() {
    //         let root = Node::Flag((), str);
    //         self.db.insert(k: K, v: V);
    //     } else {

    //     }
    // }

//     fn delete(&mut self, key: &str) {
//         // TODO
//     }
}

pub fn compact_encode(hex_array: Vec<u8>) -> Vec<u8> {
    let mut hex_data = hex_array.clone();
    let mut result = vec![];
    let len = hex_data.len();
    let term = if hex_data[len - 1] == 16 { 1 } 
        else { 0 };
    if term == 1{
        hex_data = hex_data[0..len-1].to_vec();
    }
    let oddlen = hex_data.len() % 2;
    let flags = 2 * term + oddlen;
    if oddlen == 1 {
        hex_data.insert(0, flags as u8);
    } else{
        hex_data.insert(0, 0);
        hex_data.insert(0, flags as u8);
    }
    println!("Hex Array: {:?}", hex_data);
    for i in (0..hex_data.len()).step_by(2){
        result.push(16*hex_data[i] + hex_data[i+1]);
    }
    return result
    // TODO
}

// If Leaf, ignore 16 at the end
// fn compact_decode(encoded_arr: Vec<u8>) -> Vec<u8> {
//     // TODO
// }

// fn test_compact_encode() {
//     assert_eq!(compact_decode(compact_encode(vec![1, 2, 3, 4, 5])),
//                vec![1, 2, 3, 4, 5]);
//     assert_eq!(compact_decode(compact_encode(vec![0, 1, 2, 3, 4, 5])),
//                vec![0, 1, 2, 3, 4, 5]);
//     assert_eq!(compact_decode(compact_encode(vec![0, 15, 1, 12, 11, 8, 16])),
//                vec![0, 15, 1, 12, 11, 8]);
//     assert_eq!(compact_decode(compact_encode(vec![15, 1, 12, 11, 8, 16])),
//                vec![15, 1, 12, 11, 8]);
// }

fn hash_node(node: &Node) -> String {
    let mut hasher = Sha3::sha3_256();
    match node {
        Node::Branch(branch) => {
            let mut input = String::from("branch_");
            for each in branch {
                input += &*each;
            }
            hasher.input_str(&*input);
        },
        Node::Flag((_encoded_prefix, value)) => {hasher.input_str(&*value);},
        Node::Null() => {hasher.input_str("");},
    }
    String::from("HashStart_") + &*(hasher.result_str()) + "_HashEnd"
}

pub fn string_to_hex(key: &str) -> Vec<u8>{
    let mut result = vec![];
    let hex_str = key.as_bytes().to_hex();
    for c in hex_str.chars() {
        result.push(u8::from_str_radix(&c.to_string(),16).unwrap());
    }
    println!("{:?}", result);
    return result
}