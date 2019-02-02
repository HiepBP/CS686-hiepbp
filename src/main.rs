mod merkle_patricia_trie;

include!("merkle_patricia_trie.rs");

// fn main() {
//     println!("Hello, world!");
//     println!("{:?}","0f1248".as_bytes());
//     println!("{}","b".as_bytes().to_hex());
// }

// fn main() {
//           println!("CryptoPals Rust");
//          println!("Enter Hex Number:");
//           let mut hex_input = String::new();
//           io::stdin().read_line(&mut hex_input)
//          .ok()
//           .expect("IO Error Unable to read line");
//          print!("Input was {}",hex_input);
//          let bytes = hex_input.as_bytes();
//          for b in bytes{
//               println!("{}",b);
//         }
//       }      

fn main(){
    //merkle_patricia_trie::string_to_hex("k");
    println!("{:?}",merkle_patricia_trie::compact_encode([16].to_vec()));
}
