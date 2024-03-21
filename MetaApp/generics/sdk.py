import matplotlib.pyplot as plt

filelist = ["old_sdk_online","new_sdk_online"]

step_name = [
    "user_feature_step",
    "user_fid_step",
    "sdk_pre_processor",
    "sdk_recall_step",
    "ad_sdk_filter_step",
    "predict_cache_skip_step",
    "sdk_pre_ranking_step",
    "user_cross_item_feature_step",
    "item_fid_step",
    "sdk_classify_item_step",
    "sdk_fill_game_item_fid_dispatcher_step",
    "sdk_predict_dispatcher_step",
    "sdk_model_union_dispatcher_step",
    "send_fid_step",
    "send_model_score_step",
    "sdk_post_ranking_step",
    "show_budget_control_step",
    "valuation_filter_step",
    "export_item_step",
    "sdk_post_processor"
]

resultlist = []

for filename in filelist:
    with open(filename,"r") as f:
        numlist = [0 for i in range(20)]
        count = 0
        line = f.readline()
        while line: 
            split_text = line.split('[')[-1].split(']')[0].split(" ")
            if len(split_text) == 20: 
                for step in split_text:
                    name,name_count = step.split(":")
                    numlist[step_name.index(name)] += int(name_count)
                count += 1
            line = f.readline() 

        resultlist.append([i/count for i in numlist])
    print(count)


print(resultlist[0])
print(sum(resultlist[0]))
print(resultlist[1])
print(sum(resultlist[1]))

x = [i for i in range(1,21)]
plt.figure(figsize=(30,10))

plt.bar([i-0.25 for i in range(1,21)],[i for i in resultlist[0]],width=0.5,label="old")
plt.bar([i+0.25 for i in range(1,21)],[i for i in resultlist[1]],width=0.5,label="new")
plt.xticks(x,labels=[''.join(i.split("_")[:-1]) for i in step_name],rotation=-45,size=15)
plt.yticks([i for i in range(0,55,2)],size=16)
plt.legend(fontsize=26)
plt.savefig("sdk.jpg")

# x = [i for i in range(1,20)]
# plt.figure(figsize=(30,10))

# plt.bar([i-0.25 for i in range(1,20)],[i for i in resultlist[0][1:]],width=0.5,label="old")
# plt.bar([i+0.25 for i in range(1,20)],[i for i in resultlist[1][1:]],width=0.5,label="new")
# plt.xticks(x,labels=[''.join(i.split("_")[:-1]) for i in step_name[1:]],rotation=-45,size=15)
# plt.yticks([i for i in range(0,15,1)],size=16)
# plt.legend(fontsize=26)
# plt.savefig("test1.jpg")