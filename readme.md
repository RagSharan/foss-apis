# Things needs to take care
1- log system setup
2- cross site scripting setup
3- version control setup
4- 

# Foss-API Description file

# user authentication rest apis
# user profile rest apis
# user follower and follow apis
# user comment apis
# user voting apis

# public Post apis
there are 7 post apis
# AddPost
# AddPostList
# GetPost
1- Filter needs to upgrade
2- if filter is blanck it will randomly through object needs to be fixed
filterFormate function not working to setup filter to second lavel of json
ex-  {               # here filter should be .find({"vote.upVote", 12})
    "vote": {
            "upVote": 12
        }
    }
