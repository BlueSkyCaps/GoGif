using System;
using System.Diagnostics;
using System.IO;
using System.Windows.Forms;

namespace Gif_Maker
{
    public partial class MianGui : Form
    {
        public MianGui()
        {
            InitializeComponent();
        }

        private void MianGui_Load(object sender, EventArgs e)
        {

        }

        private void MianGui_DragEnter(object sender, DragEventArgs e)
        {
            if (e.Data.GetDataPresent(DataFormats.FileDrop))
            {
                e.Effect = DragDropEffects.Link;
            }
            else
            {
                e.Effect = DragDropEffects.None;
            }
        }

        private bool ChecckArgsValidity()
        {
            if (string.IsNullOrWhiteSpace(textBox1.Text.Trim()) || string.IsNullOrWhiteSpace(textBox2.Text.Trim()) || string.IsNullOrWhiteSpace(textBox3.Text.Trim()))
            {
                MessageBox.Show("请把相关参数填写完才能生成Gif动图哦。\n最大宽度和长度：最终制作的Gif的尺寸。" +
                    "要想生成良好的效果，所有图片素材应略小于或等于输入的尺寸。有素材图片大于输入尺寸的是不行的啦！\n" +
                    "每帧间隔：每张图片停留的时长。\n" +
                    "若要翻转逆放，请勾选。", "提示", MessageBoxButtons.OK);

                return false;
            }
            try
            {
                Convert.ToInt32(textBox1.Text.Trim());
                Convert.ToInt32(textBox2.Text.Trim());
            }
            catch (Exception)
            {
                MessageBox.Show("输入的尺寸不是整数","错误",MessageBoxButtons.OK,MessageBoxIcon.Warning);
                return false;
            }
            try
            {
                Convert.ToDouble(textBox3.Text.Trim());
            }
            catch (Exception)
            {
                MessageBox.Show("输入的间隔时长不是数字", "错误", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return false;
            }
            if (Convert.ToInt32(textBox1.Text.Trim())<=5 || Convert.ToInt32(textBox2.Text.Trim())<=5)
            {
                MessageBox.Show("输入的宽或高太小啦！那么小，谁看得到？", "提示", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return false;
            }
            if (Convert.ToInt32(textBox1.Text.Trim()) >= 1024 || Convert.ToInt32(textBox2.Text.Trim()) >= 1024)
            {
                MessageBox.Show("输入的宽或高太大啦，超过1024！要生成的图片那么大，还发什么表情？", "提示", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return false;
            }
            if (Convert.ToDouble(textBox3.Text.Trim())<0.1|| Convert.ToDouble(textBox3.Text.Trim())>10)
            {
                MessageBox.Show("每帧间隔不在合理范围，输入0.1~10s都可以，间隔那么快或那么慢，谁看？", "提示", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return false;
            }
            return true;
        }

        public GoGuiCmdArgsModel CollectCmdArgs()
        {
            GoGuiCmdArgsModel args = new GoGuiCmdArgsModel
            {
                W = textBox1.Text.Trim(),
                H = textBox2.Text.Trim(),
                Dur = textBox3.Text.Trim(),
                Order = checkBox1.Checked?"1":"0",
            };
            return args;
        }

        private void MianGui_DragDrop(object sender, DragEventArgs e)
        {
            if (!ChecckArgsValidity())
            {
                return;
            };
            var inputImagesFolderA = ((Array)e.Data.GetData(DataFormats.FileDrop)).GetValue(0)?.ToString();
            if (!Directory.Exists(inputImagesFolderA))
            {
                MessageBox.Show("不存在文件夹！你确定拖动的是存放素材图片的文件夹而不是单独的文件吗？", "错误", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return;
            }
            // 整理Go程序所需命令行参数,此事件获取拖动到的素材文件夹路径。
            var args = CollectCmdArgs();
            args.InputRoot = inputImagesFolderA;
            Debug.WriteLine(args.InputRoot);

            var goProcessPath = Path.Combine(System.Windows.Forms.Application.StartupPath, "cmd", "GoGif.exe");
            Process.Start(goProcessPath, string.Join(" ", args.W, args.H, args.Dur, args.Order, args.InputRoot));
        }

        private void button1_Click(object sender, EventArgs e)
        {
            var r = MessageBox.Show("此程序作者：BlueSkyCaps(芝士为了玩|比尔小贤)，由作者本人制作且免费向外开源。" +
                "若你是从第三方购买而获取此程序，则代表你可能受骗。\n,点击取消关闭弹窗；点击确定将浏览此程序gitee仓库(也许有最新更新。这小程序没啥好更新的，说不定以后有更方便有用的功能，随缘吧。)", "说明", MessageBoxButtons.OKCancel, MessageBoxIcon.Asterisk);

            if (r == DialogResult.OK)
            {
                string giteeUrl = "https://gitee.com/BlueSkyCaps/GoGif";
                try
                {
                    System.Diagnostics.Process.Start(giteeUrl);
                }
                catch (Exception)
                {
                    try
                    {
                        System.Diagnostics.Process.Start("IEXPLORE.EXE", giteeUrl);
                    }
                    catch (Exception)
                    {
                        // noting handle.
                    }
                }
            }
        }
    }

    /// <summary>
    /// 执行GoGif程序的命令行参数
    /// </summary>
    public class GoGuiCmdArgsModel
    {
        // 输入素材文件夹路径
        public string InputRoot { get; set; }
        /// <summary>
        /// 最终Gif的宽
        /// </summary>
        public string W { get; set; }
        /// <summary>
        /// 最终Gif的高
        /// </summary>
        public string H { get; set; }
        /// <summary>
        /// 每帧(图片)的停留时长
        /// </summary>
        public string Dur { get; set; }
        /// <summary>
        /// 默认升序还是倒序生成
        /// </summary>
        public string Order { get; set; }
    }
}
