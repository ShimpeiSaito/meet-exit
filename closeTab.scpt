tell application "Google Chrome"
    beep 3
	set windowList to every tab of every window whose URL starts with "https://meet.google.com/"
	repeat with tabsByWindow in windowList
		repeat with tabItm in tabsByWindow
			delete tabItm
		end repeat
	end repeat
end tell
