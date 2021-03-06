#ifndef GLFWGO_CALLBACK_H
#define GLFWGO_CALLBACK_H

#include <GLFW/glfw3.h>

void setErrorCallback();
void clearErrorCallback();

void setFocusCallback(GLFWwindow* w);
void clearFocusCallback(GLFWwindow* w);

void setCloseCallback(GLFWwindow* w);
void clearCloseCallback(GLFWwindow* w);

void setResizeCallback(GLFWwindow* w);
void clearResizeCallback(GLFWwindow* w);

void setFramebufferResizeCallback(GLFWwindow* w);
void clearFramebufferResizeCallback(GLFWwindow* w);

void setKeyCallback(GLFWwindow* w);
void clearKeyCallback(GLFWwindow* w);

void setCharCallback(GLFWwindow* w);
void clearCharCallback(GLFWwindow* w);

void setMouseMoveCallback(GLFWwindow* w);
void clearMouseMoveCallback(GLFWwindow* w);

void setMouseButtonCallback(GLFWwindow* w);
void clearMouseButtonCallback(GLFWwindow* w);

void setScrollCallback(GLFWwindow* w);
void clearScrollCallback(GLFWwindow* w);

#endif
