import axios from 'axios';

const API_BASE_URL = process.env.REACT_APP_BACKEND_URL;

interface Project {
    id?: number;
    name: string;
}

interface Task {
    id?: number;
    name: string;
    projectId: number;
}

class ApiError extends Error {
    constructor(message: string, public status: number) {
        super(message);
        this.name = "ApiError";
    }
}

const makeApiCall = async <T>(method: "get" | "post" | "put" | "delete", url: string, data?: any): Promise<T> => {
    const config = {
        method,
        url: `${API_BASE_URL}${url}`,
        data,
    };

    try {
        const response = await axios(config);
        return response.data;
    } catch (error: any) {
        throw new ApiError(error.response?.data?.message || "An unknown error occurred", error.response?.status || 500);
    }
};

const fetchProjects = async (): Promise<Project[]> => {
    return makeApiCall<Project[]>('get', '/projects');
};

const fetchTasks = async (projectId: number): Promise<Task[]> => {
    return makeApiCall<Task[]>('get', `/projects/${projectId}/tasks`);
};

const filterTasksByName = async (projectId: number, taskName: string): Promise<Task[]> => {
    const tasks = await fetchTasks(projectId);
    return tasks.filter(task => task.name.toLowerCase().includes(taskName.toLowerCase()));
};

const createProject = async (project: Project): Promise<Project> => {
    return makeApiCall<Project>('post', '/projects', project);
};

const updateProject = async (project: Project): Promise<Project> => {
    return makeApiCall<Project>('put', `/projects/${project.id}`, project);
};

const deleteProject = async (id: number): Promise<void> => {
    await makeApiCall<void>('delete', `/projects/${id}`);
};

const createTask = async (task: Task): Promise<Task> => {
    return makeApiCall<Task>('post', '/tasks', task);
};

const updateTask = async (task: Task): Promise<Task> => {
    return makeApiCall<Task>('put', `/tasks/${task.id}`, task);
};

const deleteTask = async (id: number): Promise<void> => {
    await makeApiCall<void>('delete', `/tasks/${id}`);
};

const bulkDeleteTasks = async (projectId: number): Promise<void> => {
    const tasks = await fetchTasks(projectId);
    const deletePromises = tasks.map(task => deleteTask(task.id!));
    await Promise.all(deletePromises);
};

document.addEventListener("DOMContentLoaded", async () => {
    try {
        const projects = await fetchProjects();
        console.log(projects);
    } catch (error) {
        console.error("Failed to fetch projects", error instanceof Error ? error.message : error);
    }
});

document.getElementById("create-project")?.addEventListener("click", async () => {
    const projectNameElement = document.getElementById("project-name") as HTMLInputElement;
    if (!projectNameElement) return;
    const projectName = projectNameElement.value;
    try {
        const newProject = await createProject({ name: projectName });
        console.log(newProject);
    } catch (error) {
        console.error("Failed to create a project", error instanceof Error ? error.message : error);
    }
});