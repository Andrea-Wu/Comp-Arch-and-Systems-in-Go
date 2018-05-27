import io 
import sys #command line args

class Hashtable:
    def __init__(self, size):
        self.size = size
        self.table = {-1: "null"}

    def hashFunction(self, myInt):
        return myInt % self.size

    def insert(self, data):
        newNode = Node(data, None)

        bucket = self.hashFunction(data)

        if self.search(data):
            print("duplicate: did not insert " + str(data))
            return

        try: #if the key exists in the dict
            temp = self.table[bucket]
            self.table[bucket] = newNode
            newNode.nextNode = temp 
        except: #else
            self.table[bucket] = newNode

        print("inserted " + str(data))

            
    def search(self, find_this):
        bucket = self.hashFunction(find_this)

        try:        
            head = self.table[bucket]
            while head != None:
                if head.data == find_this:
                    print(str(find_this) + " was found")
                    return True
                head = head.nextNode
                
            print(str(find_this) + " is absent")

        except:
            print(str(find_this) + " is absent")
            return False

        return False
                
        

class Node:
    def __init__(self, data, nextNode):
        self.data = data
        self.nextNode = nextNode

    
def main():
    myHash = Hashtable(1000)
    myHash.insert(10)

    #open text file for consumption
    fp = open(sys.argv[1], "r")
    
    char = '\0'
    for line in fp:
        char = line[0]
        i = 1
        while line[i] == ' ':
            i = i + 1
        
        num = int(line[i:])

        if char == 'i':
            myHash.insert(num)

        elif char == 's':
            myHash.search(num)
    
        else:
            print("err")
     

    fp.close()

if __name__ == "__main__":
    main()
