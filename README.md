# PromethoniXTrie

**Promethon:** New way to look at the next generation of Internet!

Ethereum uses a world state to hold states of each account.
But the size of this world state is growing rapidly and there hasn't been any efficient solution for shrinking it down.
World state is implemented by using [Merkle Patricia Trie](https://ethereum.org/en/developers/docs/data-structures-and-encoding/patricia-merkle-trie/) but today, we introduce **PromethoniXTrie**.
PromethoniXTrie acts as a bridge between [Merkle Patricia Trie](https://ethereum.org/en/developers/docs/data-structures-and-encoding/patricia-merkle-trie/) and [Red-Black Tree](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree).
The main idea is to **keep an extra value to validate** that the account should still exist on the network or not.
But due to the large volume of data, the challenge that has always existed is how to perform this validation in a short time and remove invalid accounts.

> A [Red-Black Tree](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree) is a specialised binary search tree data structure noted for fast storage and retrieval of ordered information, and a guarantee that operations will complete within a known time. Also, guarantees searching in Big O time of O(logn), where n is the number of entries (or keys) in the tree. The insert and delete operations, along with the tree rearrangement and recoloring, are also performed in O(logn) time.

Therefore, according to the description of the Red-Black Tree, if we keep a Red-Black Tree next to our Merkle Patricia Trie and can establish a relationship between them so that the Red-Black Tree is always sorted based on the additional value we keep, we can always find the leftmost node of Red-Black Tree in Big O time of **O(1)**.

Now the question that arises is, **what is this extra value and how does it perform the validation?**

To be honest, this question can have many different answers, each with pros and cons. But what we all know is that eventually we want to reduce the volume of data, so we can't store all the data all the time.

- An initial idea: we can consider this "extra value" as the block time of the last change or use in the contract. For example, if a year has passed since the last change, we can delete that account.
- Another idea is to charge a fee per certain amount of volume. Whenever this amount is used up and never charged again, delete that account.

The **PromethoniXTrie** project implements the basis of the second idea. It creates a Merkle Patricia Trie and a Red-Black Tree and establishes the desired connection between them and then tests the performance of the **Promethon** idea. You can read more about the algorithm [here](https://github.com/Promethon/PromethoniXTrie/blob/main/Promethon.pdf).

Our initial tests have been positive and promising. We are currently implementing this idea on the [go-ethereum](https://github.com/ethereum/go-ethereum) source code and we hope to announce its completion soon!
