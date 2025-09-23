"use client";

import type React from "react";

import { useEffect, useState } from "react";
import Link from "next/link";
import { Input } from "@/components/ui/input";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Separator } from "@/components/ui/separator";
import {
  Hash,
  Plus,
  Settings,
  Mic,
  Headphones,
  Phone,
  Video,
  UserPlus,
  Search,
  Smile,
  Gift,
  FileImage,
  MoreHorizontal,
  Volume2,
  MicOff,
  LogOut,
} from "lucide-react";
import { theme } from "@/lib/theme";

import { useQuery } from "@tanstack/react-query";

export default function DashboardPage() {
  const [selectedChannel, setSelectedChannel] = useState("general");
  const [message, setMessage] = useState("");
  const [isMuted, setIsMuted] = useState(false);

  const { data, error, isLoading } = useQuery({
    queryKey: ["servers"],
    queryFn: async () => {
      return (await (await fetch("/api/server")).json()).result ?? null;
    },
  });

  useEffect(() => {
    console.log(`data: ${JSON.stringify(data, null, 2)}`);
  }, [data]);
  useEffect(() => {
    console.log(`error: ${error}`);
  }, [error]);
  useEffect(() => {
    console.log(`isLoading: ${isLoading}`);
  }, [isLoading]);

  const servers = [
    { id: "1", name: "My Server", avatar: "MS", color: "bg-pink-600" },
    { id: "2", name: "Gaming Hub", avatar: "GH", color: "bg-blue-600" },
    { id: "3", name: "Study Group", avatar: "SG", color: "bg-green-600" },
  ];

  const channels = [
    { id: "general", name: "general", type: "text" },
    { id: "random", name: "random", type: "text" },
    { id: "gaming", name: "gaming", type: "text" },
    { id: "voice1", name: "General Voice", type: "voice" },
    { id: "voice2", name: "Gaming Voice", type: "voice" },
  ];

  const messages = [
    {
      id: 1,
      user: "Alice",
      avatar: "A",
      content: "Hey everyone! How's it going?",
      timestamp: "Today at 2:30 PM",
      color: "bg-purple-600",
    },
    {
      id: 2,
      user: "Bob",
      avatar: "B",
      content: "Pretty good! Just working on some code.",
      timestamp: "Today at 2:32 PM",
      color: "bg-blue-600",
    },
    {
      id: 3,
      user: "Charlie",
      avatar: "C",
      content: "Anyone want to play some games later?",
      timestamp: "Today at 2:35 PM",
      color: "bg-green-600",
    },
  ];

  const onlineUsers = [
    {
      id: 1,
      name: "Alice",
      status: "online",
      avatar: "A",
      color: "bg-purple-600",
    },
    { id: 2, name: "Bob", status: "idle", avatar: "B", color: "bg-blue-600" },
    {
      id: 3,
      name: "Charlie",
      status: "dnd",
      avatar: "C",
      color: "bg-green-600",
    },
    {
      id: 4,
      name: "Diana",
      status: "offline",
      avatar: "D",
      color: "bg-red-600",
    },
  ];

  const handleSendMessage = (e: React.FormEvent) => {
    e.preventDefault();
    if (message.trim()) {
      console.log("Sending message:", message);
      setMessage("");
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case "online":
        return "bg-green-500";
      case "idle":
        return "bg-yellow-500";
      case "dnd":
        return "bg-red-500";
      default:
        return "bg-gray-500";
    }
  };

  return (
    <div className={`${theme.classes.background} h-screen flex`}>
      {/* Server Sidebar */}
      <div className="w-16 bg-gray-900 flex flex-col items-center py-3 space-y-2">
        {servers.map((server) => (
          <div
            key={server.id}
            className={`w-12 h-12 ${server.color} rounded-2xl hover:rounded-xl transition-all duration-200 flex items-center justify-center cursor-pointer group relative`}
          >
            <span className={`${theme.colors.text.primary} font-bold text-sm`}>
              {server.avatar}
            </span>
            <div className="absolute left-full ml-2 px-2 py-1 bg-black text-white text-xs rounded opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap z-10">
              {server.name}
            </div>
          </div>
        ))}
        <div className="w-12 h-12 bg-gray-700 hover:bg-pink-600 rounded-2xl hover:rounded-xl transition-all duration-200 flex items-center justify-center cursor-pointer">
          <Plus className="w-6 h-6 text-green-400" />
        </div>
      </div>

      {/* Channel Sidebar */}
      <div className="w-60 bg-gray-800 flex flex-col">
        {/* Server Header */}
        <div className="h-12 px-4 flex items-center justify-between border-b border-gray-700 shadow-sm">
          <h2 className={`${theme.colors.text.primary} font-semibold`}>
            My Server
          </h2>
          <MoreHorizontal
            className={`w-4 h-4 ${theme.colors.text.secondary} cursor-pointer hover:text-white`}
          />
        </div>

        {/* Channels */}
        <ScrollArea className="flex-1 px-2 py-3">
          <div className="space-y-1">
            <div className="flex items-center justify-between px-2 py-1">
              <span
                className={`${theme.colors.text.secondary} text-xs font-semibold uppercase tracking-wide`}
              >
                Text Channels
              </span>
              <Plus
                className={`w-4 h-4 ${theme.colors.text.secondary} cursor-pointer hover:text-white`}
              />
            </div>

            {channels
              .filter((c) => c.type === "text")
              .map((channel) => (
                <div
                  key={channel.id}
                  onClick={() => setSelectedChannel(channel.id)}
                  className={`flex items-center px-2 py-1 rounded cursor-pointer group ${
                    selectedChannel === channel.id
                      ? "bg-gray-700 text-white"
                      : `${theme.colors.text.secondary} hover:bg-gray-700 hover:text-white`
                  }`}
                >
                  <Hash className="w-4 h-4 mr-2" />
                  <span className="text-sm">{channel.name}</span>
                </div>
              ))}

            <div className="flex items-center justify-between px-2 py-1 mt-4">
              <span
                className={`${theme.colors.text.secondary} text-xs font-semibold uppercase tracking-wide`}
              >
                Voice Channels
              </span>
              <Plus
                className={`w-4 h-4 ${theme.colors.text.secondary} cursor-pointer hover:text-white`}
              />
            </div>

            {channels
              .filter((c) => c.type === "voice")
              .map((channel) => (
                <div
                  key={channel.id}
                  className={`flex items-center px-2 py-1 rounded cursor-pointer group ${theme.colors.text.secondary} hover:bg-gray-700 hover:text-white`}
                >
                  <Volume2 className="w-4 h-4 mr-2" />
                  <span className="text-sm">{channel.name}</span>
                </div>
              ))}
          </div>
        </ScrollArea>

        {/* User Panel */}
        <div className="h-14 bg-gray-900 px-2 flex items-center justify-between">
          <div className="flex items-center space-x-2">
            <Avatar className="w-8 h-8">
              <AvatarFallback className="bg-pink-600 text-white text-sm">
                U
              </AvatarFallback>
            </Avatar>
            <div className="flex-1 min-w-0">
              <p
                className={`${theme.colors.text.primary} text-sm font-medium truncate`}
              >
                You
              </p>
              <p className={`${theme.colors.text.secondary} text-xs truncate`}>
                Online
              </p>
            </div>
          </div>
          <div className="flex items-center space-x-1">
            <button
              onClick={() => setIsMuted(!isMuted)}
              className={`p-1 rounded hover:bg-gray-700 ${isMuted ? "text-red-400" : theme.colors.text.secondary}`}
            >
              {isMuted ? (
                <MicOff className="w-4 h-4" />
              ) : (
                <Mic className="w-4 h-4" />
              )}
            </button>
            <button
              className={`p-1 rounded hover:bg-gray-700 ${theme.colors.text.secondary}`}
            >
              <Headphones className="w-4 h-4" />
            </button>
            <button
              className={`p-1 rounded hover:bg-gray-700 ${theme.colors.text.secondary}`}
            >
              <Settings className="w-4 h-4" />
            </button>
          </div>
        </div>
      </div>

      {/* Main Content */}
      <div className="flex-1 flex flex-col bg-gray-700">
        {/* Channel Header */}
        <div className="h-12 px-4 flex items-center justify-between border-b border-gray-600 bg-gray-700">
          <div className="flex items-center space-x-2">
            <Hash className={`w-5 h-5 ${theme.colors.text.secondary}`} />
            <span className={`${theme.colors.text.primary} font-semibold`}>
              {selectedChannel}
            </span>
            <Separator orientation="vertical" className="h-6 bg-gray-600" />
            <span className={`${theme.colors.text.secondary} text-sm`}>
              Welcome to #{selectedChannel}!
            </span>
          </div>
          <div className="flex items-center space-x-2">
            <Phone
              className={`w-5 h-5 ${theme.colors.text.secondary} cursor-pointer hover:text-white`}
            />
            <Video
              className={`w-5 h-5 ${theme.colors.text.secondary} cursor-pointer hover:text-white`}
            />
            <UserPlus
              className={`w-5 h-5 ${theme.colors.text.secondary} cursor-pointer hover:text-white`}
            />
            <Search
              className={`w-5 h-5 ${theme.colors.text.secondary} cursor-pointer hover:text-white`}
            />
            <Link href="/login">
              <LogOut
                className={`w-5 h-5 ${theme.colors.text.secondary} cursor-pointer hover:text-white`}
              />
            </Link>
          </div>
        </div>

        {/* Messages */}
        <ScrollArea className="flex-1 px-4 py-4">
          <div className="space-y-4">
            {messages.map((msg) => (
              <div
                key={msg.id}
                className="flex items-start space-x-3 hover:bg-gray-600/30 px-2 py-1 rounded"
              >
                <Avatar className="w-10 h-10 mt-1">
                  <AvatarFallback className={`${msg.color} text-white`}>
                    {msg.avatar}
                  </AvatarFallback>
                </Avatar>
                <div className="flex-1 min-w-0">
                  <div className="flex items-baseline space-x-2">
                    <span
                      className={`${theme.colors.text.primary} font-medium text-sm`}
                    >
                      {msg.user}
                    </span>
                    <span className={`${theme.colors.text.secondary} text-xs`}>
                      {msg.timestamp}
                    </span>
                  </div>
                  <p className={`${theme.colors.text.primary} text-sm mt-1`}>
                    {msg.content}
                  </p>
                </div>
              </div>
            ))}
          </div>
        </ScrollArea>

        {/* Message Input */}
        <div className="p-4">
          <form onSubmit={handleSendMessage} className="relative">
            <Input
              value={message}
              onChange={(e) => setMessage(e.target.value)}
              placeholder={`Message #${selectedChannel}`}
              className="bg-gray-600 border-none text-white placeholder:text-gray-400 pr-12 py-3 rounded-lg focus-visible:ring-1 focus-visible:ring-pink-600"
            />
            <div className="absolute right-3 top-1/2 -translate-y-1/2 flex items-center space-x-1">
              <button
                type="button"
                className={`p-1 rounded hover:bg-gray-500 ${theme.colors.text.secondary}`}
              >
                <Gift className="w-4 h-4" />
              </button>
              <button
                type="button"
                className={`p-1 rounded hover:bg-gray-500 ${theme.colors.text.secondary}`}
              >
                <FileImage className="w-4 h-4" />
              </button>
              <button
                type="button"
                className={`p-1 rounded hover:bg-gray-500 ${theme.colors.text.secondary}`}
              >
                <Smile className="w-4 h-4" />
              </button>
            </div>
          </form>
        </div>
      </div>

      {/* Members Sidebar */}
      <div className="w-60 bg-gray-800">
        <div className="p-4">
          <h3
            className={`${theme.colors.text.primary} font-semibold text-sm mb-3`}
          >
            Members — {onlineUsers.length}
          </h3>
          <div className="space-y-1">
            {onlineUsers.map((user) => (
              <div
                key={user.id}
                className="flex items-center space-x-3 px-2 py-1 rounded hover:bg-gray-700 cursor-pointer"
              >
                <div className="relative">
                  <Avatar className="w-8 h-8">
                    <AvatarFallback
                      className={`${user.color} text-white text-sm`}
                    >
                      {user.avatar}
                    </AvatarFallback>
                  </Avatar>
                  <div
                    className={`absolute -bottom-1 -right-1 w-3 h-3 ${getStatusColor(user.status)} rounded-full border-2 border-gray-800`}
                  />
                </div>
                <span className={`${theme.colors.text.primary} text-sm`}>
                  {user.name}
                </span>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}
