## Clipboard History Project

Imagine you're working on a research project that involves gathering information from multiple sources. You need to copy various text snippets, URLs, and notes from different websites, documents, and emails. Usually, this process involves repeatedly switching between applications to recopy the same content, which can be quite cumbersome. 

### That's where clip-save comes to save you.You can now track upto past 20 clipboard contents at a time.
The project consists of two main components:

### 1. Server

The server is responsible for monitoring the clipboard and storing its content in text files. It performs the following tasks:

- Continuously watches the clipboard for any changes.
- Detects when new text is copied to the clipboard.
- Stores the content of the clipboard in text files, creating a history of copied text.
- Avoids duplicating content if it has been previously copied.

### 2. Display Service

The display service is a graphical user interface that provides a user-friendly way to view and interact with the stored clipboard content. It offers the following features:

- Displays a list of the stored clipboard entries.
- Allows users to navigate through the clipboard history.
- Select a specific clipboard entry to copy its content back to the clipboard.
- Enables efficient management of copied text by providing a user interface for browsing and selecting items from the history.

## How to Run the Project

To run the Clipboard History project, follow these steps:

### Starting the Server:

1. Navigate to the server directory:
   ```shell
   cd server
   ```

2. Build the server executable:
   ```shell
   go build ./main.go
   ```

3. Run the server:
   ```shell
   ./main
   ```

The server will begin monitoring the clipboard, saving copied content in text files to create a clipboard history.

### Starting the Display Service:

1. Navigate to the display directory:
   ```shell
   cd display
   ```

2. Build the display service executable:
   ```shell
   go build ./main.go
   ```

3. Run the display service:
   ```shell
   ./main
   ```

The display service opens a graphical window where you can access and interact with the clipboard history.

## Usage

- The server continuously monitors and stores clipboard changes, building a history of copied text.
- The display service provides an easy-to-use graphical interface for viewing, selecting, and managing clipboard history entries.

## Challenges:

### Clipboard Monitoring: 
Implementing a clipboard monitoring service can be challenging. This component needs to continuously track changes in the clipboard and respond to events when new content is copied.

### File Management: 
The server component is responsible for writing and managing text files for clipboard entries. Ensuring that content is stored efficiently and preventing duplicate entries requires careful file handling.

### UI Development: 
Creating a user interface for managing clipboard history can be a complex task. You need to design a clean and intuitive interface, handle user interactions, and ensure that the application is responsive.

Overall, this project involves both low-level system interaction (clipboard monitoring and file I/O) and high-level user interface design, making it an interesting blend of system-level and application-level development. It's a great example of how software can solve practical problems in a user-friendly manner.

![Screenshot 2023-10-31 at 1.20.29 PM.png](resources%2FScreenshot%202023-10-31%20at%201.20.29%20PM.png)
![Screenshot 2023-10-31 at 1.21.12 PM.png](resources%2FScreenshot%202023-10-31%20at%201.21.12%20PM.png)
![Screenshot 2023-10-31 at 1.22.22 PM.png](resources%2FScreenshot%202023-10-31%20at%201.22.22%20PM.png)


