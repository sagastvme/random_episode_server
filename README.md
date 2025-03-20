# Video Streaming App

This Go project allows you to stream random episodes from a collection of videos stored in a "videos" folder. Every time you open the website, it picks a random episode and starts playing at a random minute, just like when you turn on the TV and find an episode already playing!

## Features
- **Random Episode Selection**: Every time you access the site, a random episode is chosen from your video collection.
- **Random Start Time**: The episode starts at a random minute, adding an element of surprise each time you watch.
- **Docker Support**: The project includes a Dockerfile for easy containerization.

## Installation

### Option 1: Running with Docker (Recommended)

1. **Clone the repository**:

   ```bash
   git clone <repository-url>
   cd <project-folder>
   ```

2. **Create the `videos` folder**:
   - In the root directory of the project, create a folder named `videos`.
   - Add your video files (in formats like `.mp4`, `.avi`, etc.) to this folder.

   ```bash
   mkdir videos
   ```

3. **Build and run the Docker container**:
   - The project includes a `Dockerfile` for easy setup. Build the Docker image using the following command:

   ```bash
   docker build -t random_episode_server .
   ```

   - After building the image, run the container:

   ```bash
   docker run -p 8080:8080 -v $(pwd)/videos:/app/videos random_episode_server
   ```

   This will start a local web server and make the app accessible at `http://localhost:8080`.


## How it Works

When you visit the website:
1. The app scans the `videos` folder for video files.
2. A random video is selected from the available files.
3. A random start time in the video is chosen.
4. The video plays from that random point, simulating a TV-like experience.

### Example:

- If you have `video1.mp4` and `video2.mp4` in the `videos` folder, the website will randomly select one of them and start playing at a random point.

## Requirements

- Docker (optional, but recommended for ease of setup).
- Go 1.16 or higher (if not using Docker).
- A collection of video files stored in the `videos` folder.