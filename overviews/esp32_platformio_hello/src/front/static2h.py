import os, re

PATH   = "static/"
H_FILE = "static.h"


filenames    = os.listdir(PATH)
contents_str = ""
files_str    = ""

print("")
for filename in filenames:
    content_type = ""
    if not os.path.isfile(PATH + filename):
        if filename != "." and filename != "..":
            print("WARNING: directory " + PATH + filename + " is omitted")
        continue
    elif filename.endswith(".htm") or filename.endswith(".html"):
        content_type = "text/html"
    elif filename.endswith(".js"):
        content_type = "text/javascript" # ??? "application/javascript"
    elif filename.endswith(".css"):
        content_type = "text/css"
    elif filename.endswith(".png"):
        content_type = "image/png"
    elif filename.endswith(".jpg") or filename.endswith(".jpeg"):
        content_type = "image/jpeg"
    elif filename.endswith(".svg"):
        content_type = "image/svg+xml"
    else:
        print("WARNING: file " + PATH + filename + " is omitted")
        continue

    try:
        f = open(PATH + filename, 'rb')
        content = f.read()
        f.close()
    except:
        print("ERROR: can't read file " + PATH + filename)
        continue

    print("processed file " + PATH + filename)
    
    var_name      = re.sub('[^0-9a-zA-Z_]', '_', filename)
    contents_str += "uint8_t " + var_name + "[] = {" + ",".join([hex(ch) for ch in content]) + "};\n"
    files_str    += "    {\"/" + filename + "\", \"" + content_type + "\", (uint8_t*)&" + var_name + ", sizeof(" + var_name + ")},\n"
    if filename == "index.html" or filename == "index.htm":
        files_str += "    {\"/\", \"" + content_type + "\", (uint8_t*)&" + var_name + ", sizeof(" + var_name + ")},\n"

    
try:
    f = open(H_FILE, 'w')
    f.write("#pragma once\n\n" + \
            "#include <stddef.h>\n" + \
            "#include \"static_files.h\"\n\n" + \
            contents_str + "\n" + \
            "static_file files[] = {\n" + files_str + "};\n\n" + \
            "size_t FILES_LENGTH = sizeof(files) / sizeof(static_file);\n")
    f.close()
    print(H_FILE + " is written")
except:
    print("ERROR: can't write file " + H_FILE)
    
